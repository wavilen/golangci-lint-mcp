#!/usr/bin/env bash
# compare-e2e-runs.sh — Compare E2E test results between two version specs on the e2e-artifacts branch.
# Reads result JSONs via git show (no checkout needed), diffs all D-05 fields,
# and flags ANY change — improvements, regressions, and status changes.
set -euo pipefail

# --- Dependency check ---
if ! command -v jq >/dev/null 2>&1; then
    echo "Error: jq is required but not installed." >&2
    echo "Install with: apt install jq / brew install jq" >&2
    exit 1
fi

# --- Argument validation ---
if [ $# -ne 2 ]; then
    echo "Usage: compare-e2e-runs.sh <old_version_spec> <new_version_spec>"
    echo "Example: compare-e2e-runs.sh v1.4.0/2026-05-01 v1.3.0/2026-04-27"
    exit 1
fi

OLD_SPEC="$1"
NEW_SPEC="$2"

# --- Fetch e2e-artifacts branch ---
git fetch origin e2e-artifacts 2>/dev/null || true

BRANCH="e2e-artifacts"

# --- Verify both version specs exist ---
verify_version_spec() {
    local spec="$1"
    local label="$2"
    if ! git ls-tree "$BRANCH" "${spec}/" --name-only >/dev/null 2>&1; then
        echo "Error: version spec '${spec}' not found on ${BRANCH} branch" >&2
        exit 1
    fi
}

verify_version_spec "$OLD_SPEC" "old"
verify_version_spec "$NEW_SPEC" "new"

# --- Discover result files in both version specs ---
OLD_FILES=$(git ls-tree "$BRANCH" "${OLD_SPEC}/results/" --name-only 2>/dev/null | sed "s|${OLD_SPEC}/results/||" || true)
NEW_FILES=$(git ls-tree "$BRANCH" "${NEW_SPEC}/results/" --name-only 2>/dev/null | sed "s|${NEW_SPEC}/results/||" || true)

if [ -z "$OLD_FILES" ] && [ -z "$NEW_FILES" ]; then
    echo "No result files found in either version spec."
    exit 0
fi

# --- Header ---
echo "E2E Comparison: ${OLD_SPEC} → ${NEW_SPEC}"
echo ""

# --- State tracking for summary ---
TOTAL_COMPARED=0
TOTAL_IMPROVED=0
TOTAL_REGRESSED=0
TOTAL_CHANGED=0
TOTAL_UNCHANGED=0
ONLY_OLD=""
ONLY_NEW=""

# --- Field definitions (D-05) ---
# Format: field_name type (int|float|bool|string)
FIELDS="BeforeIssues int
AfterIssues int
IssueReduction float
Pass bool
ToolCalls int
NolintCount int
ConfigModified bool"

# --- Compare a numeric (int/float) field ---
compare_numeric() {
    local label="$1" old="$2" new="$3"
    # Use awk for float comparison to avoid bash limitations
    if [ "$old" = "$new" ]; then
        echo "  ${label}:$(printf '%*s' $((20 - ${#label})) '') ${old}  (no change)"
        return 0
    fi
    # Determine direction: lower is better for issues/nolint, higher is better for reduction/toolcalls
    # For generic comparison, just show the delta direction
    local direction=""
    # For issues/counts: lower is better
    case "$label" in
        BeforeIssues|AfterIssues|NolintCount|ToolCalls)
            if awk "BEGIN { exit ($new < $old) ? 0 : 1 }"; then
                direction="✓ improved"
                TOTAL_IMPROVED=$((TOTAL_IMPROVED + 1))
            else
                direction="✗ regressed"
                TOTAL_REGRESSED=$((TOTAL_REGRESSED + 1))
            fi
            ;;
        IssueReduction)
            if awk "BEGIN { exit ($new > $old) ? 0 : 1 }"; then
                direction="✓ improved"
                TOTAL_IMPROVED=$((TOTAL_IMPROVED + 1))
            else
                direction="✗ regressed"
                TOTAL_REGRESSED=$((TOTAL_REGRESSED + 1))
            fi
            ;;
        *)
            direction="✗ CHANGED"
            TOTAL_CHANGED=$((TOTAL_CHANGED + 1))
            ;;
    esac
    echo "  ${label}:$(printf '%*s' $((20 - ${#label})) '') ${old} → ${new}  ${direction}"
    return 1
}

# --- Compare a boolean field ---
compare_bool() {
    local label="$1" old="$2" new="$3"
    if [ "$old" = "$new" ]; then
        echo "  ${label}:$(printf '%*s' $((20 - ${#label})) '') ${old}  (no change)"
        return 0
    fi
    local direction=""
    case "$label" in
        Pass|ConfigModified)
            if [ "$new" = "true" ] && [ "$label" = "Pass" ]; then
                direction="✓ improved"
                TOTAL_IMPROVED=$((TOTAL_IMPROVED + 1))
            elif [ "$new" = "false" ] && [ "$label" = "Pass" ]; then
                direction="✗ regressed"
                TOTAL_REGRESSED=$((TOTAL_REGRESSED + 1))
            else
                direction="✗ CHANGED"
                TOTAL_CHANGED=$((TOTAL_CHANGED + 1))
            fi
            ;;
        *)
            direction="✗ CHANGED"
            TOTAL_CHANGED=$((TOTAL_CHANGED + 1))
            ;;
    esac
    echo "  ${label}:$(printf '%*s' $((20 - ${#label})) '') ${old} → ${new}  ${direction}"
    return 1
}

# --- Extract field value from a result JSON ---
extract_field() {
    local content="$1"
    local field="$2"
    echo "$content" | jq -r --arg f "$field" '.[$f]'
}

# --- Compare two result JSONs ---
compare_results() {
    local filename="$1"
    local old_content="$2"
    local new_content="$3"
    local has_changes=0

    # Parse test case and model from filename (e.g., simple-glm-5-turbo-result.json)
    local testcase
    testcase=$(echo "$filename" | sed 's/-result\.json$//' | sed 's/-[^-]*$//')
    local model
    model=$(echo "$filename" | sed 's/-result\.json$//' | sed "s/^${testcase}-//")

    echo "${testcase}/${model}:"

    # Compare standard D-05 fields
    while IFS= read -r line; do
        field=$(echo "$line" | awk '{print $1}')
        ftype=$(echo "$line" | awk '{print $2}')

        old_val=$(extract_field "$old_content" "$field")
        new_val=$(extract_field "$new_content" "$field")

        case "$ftype" in
            int|float)
                if ! compare_numeric "$field" "$old_val" "$new_val"; then
                    has_changes=1
                fi
                ;;
            bool)
                if ! compare_bool "$field" "$old_val" "$new_val"; then
                    has_changes=1
                fi
                ;;
        esac
    done <<< "$FIELDS"

    # Compare ConfigDiff (string field — show if changed)
    local old_configdiff new_configdiff
    old_configdiff=$(extract_field "$old_content" "ConfigDiff")
    new_configdiff=$(extract_field "$new_content" "ConfigDiff")
    # Handle null values from jq
    if [ "$old_configdiff" = "null" ]; then old_configdiff="(none)"; fi
    if [ "$new_configdiff" = "null" ]; then new_configdiff="(none)"; fi

    if [ "$old_configdiff" = "$new_configdiff" ]; then
        echo "  ConfigDiff:$(printf '%*s' $((17)) '') (no change)"
    else
        echo "  ConfigDiff:$(printf '%*s' $((17)) '') ${old_configdiff} → ${new_configdiff}  ✗ CHANGED"
        has_changes=1
        TOTAL_CHANGED=$((TOTAL_CHANGED + 1))
    fi

    echo ""
    TOTAL_COMPARED=$((TOTAL_COMPARED + 1))

    if [ "$has_changes" -eq 0 ]; then
        TOTAL_UNCHANGED=$((TOTAL_UNCHANGED + 1))
    fi
}

# --- Build combined file list ---
ALL_FILES=$(echo -e "${OLD_FILES}\n${NEW_FILES}" | sort -u)

# --- Process each file ---
for f in $ALL_FILES; do
    # Skip empty lines
    [ -z "$f" ] && continue

    old_has=false
    new_has=false

    # Check if file exists in old spec
    if echo "$OLD_FILES" | grep -qxF "$f"; then
        old_has=true
    fi
    # Check if file exists in new spec
    if echo "$NEW_FILES" | grep -qxF "$f"; then
        new_has=true
    fi

    if [ "$old_has" = "true" ] && [ "$new_has" = "true" ]; then
        # Both exist — compare them
        old_content=$(git show "${BRANCH}:${OLD_SPEC}/results/${f}" 2>/dev/null)
        new_content=$(git show "${BRANCH}:${NEW_SPEC}/results/${f}" 2>/dev/null)
        compare_results "$f" "$old_content" "$new_content"
    elif [ "$old_has" = "true" ]; then
        # Only in old — removed
        testcase=$(echo "$f" | sed 's/-result\.json$//' | sed 's/-[^-]*$//')
        model=$(echo "$f" | sed 's/-result\.json$//' | sed "s/^${testcase}-//")
        echo "${testcase}/${model}: REMOVED (exists in ${OLD_SPEC} but not ${NEW_SPEC})"
        echo ""
        if [ -n "$ONLY_OLD" ]; then ONLY_OLD="${ONLY_OLD}, "; fi
        ONLY_OLD="${ONLY_OLD}${testcase}/${model}"
    else
        # Only in new — new test
        testcase=$(echo "$f" | sed 's/-result\.json$//' | sed 's/-[^-]*$//')
        model=$(echo "$f" | sed 's/-result\.json$//' | sed "s/^${testcase}-//")
        echo "${testcase}/${model}: NEW (exists in ${NEW_SPEC} but not ${OLD_SPEC})"
        echo ""
        if [ -n "$ONLY_NEW" ]; then ONLY_NEW="${ONLY_NEW}, "; fi
        ONLY_NEW="${ONLY_NEW}${testcase}/${model}"
    fi
done

# --- Summary ---
TOTAL_FLAGS=$((TOTAL_IMPROVED + TOTAL_REGRESSED + TOTAL_CHANGED))
echo "Summary: ${TOTAL_COMPARED} test combos compared, ${TOTAL_FLAGS} changes flagged (${TOTAL_IMPROVED} improvements, ${TOTAL_REGRESSED} regressions, ${TOTAL_CHANGED} config/other changes)"
if [ -n "$ONLY_NEW" ]; then
    echo "New tests: ${ONLY_NEW}"
fi
if [ -n "$ONLY_OLD" ]; then
    echo "Removed tests: ${ONLY_OLD}"
fi
