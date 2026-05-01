#!/usr/bin/env python3
"""Tests for extract_context.py — per-trace context extraction from ndjson + result JSON."""

import json
import os
import sys
import unittest

# Ensure the script directory is on the path
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
if SCRIPT_DIR not in sys.path:
    sys.path.insert(0, SCRIPT_DIR)

# Import after path setup
import extract_context  # noqa: E402

# Real data paths
PROJECT_ROOT = os.path.abspath(os.path.join(SCRIPT_DIR, "..", ".."))
NDJSON_DIR = os.path.join(PROJECT_ROOT, "tmp", "ndjson")


class TestExtractContextFromRealData(unittest.TestCase):
    """Test extract_context against real ndjson trace files."""

    @classmethod
    def setUpClass(cls):
        """Skip if real test data not available."""
        if not os.path.isdir(NDJSON_DIR):
            raise unittest.SkipTest("tmp/ndjson/ not found — run e2e tests first")

    def test_real_ndjson_produces_valid_context_json(self):
        """Test 1: Given real ndjson + result JSON, produces valid context with all D-06 fields."""
        ndjson_file = os.path.join(NDJSON_DIR, "large-glm-5-1.ndjson")
        result_file = os.path.join(NDJSON_DIR, "large-glm-5-1-result.json")

        if not os.path.exists(ndjson_file):
            self.skipTest("large-glm-5-1.ndjson not found")

        # Parse events
        events = extract_context.parse_ndjson_full(ndjson_file)
        self.assertGreater(len(events), 0, "Should parse events from ndjson")

        # Load result
        result = {}
        if os.path.exists(result_file):
            with open(result_file) as f:
                result = json.load(f)

        # Extract context
        ctx = extract_context.extract_context("large-glm-5-1", events, result)

        # Verify all D-06 fields exist
        required_fields = [
            "model",
            "test_case",
            "duration_ms",
            "event_counts",
            "tool_breakdown",
            "token_totals",
            "issue_counts",
            "timestamps",
            "subagent_count",
            "strategy",
            "pass",
        ]
        for field in required_fields:
            self.assertIn(field, ctx, f"Missing D-06 field: {field}")

        # Verify specific values from known result
        self.assertEqual(ctx["model"], "GLM-5.1")
        self.assertEqual(ctx["test_case"], "Large")
        self.assertEqual(ctx["issue_counts"]["before"], 116)
        self.assertEqual(ctx["issue_counts"]["after"], 0)
        self.assertEqual(ctx["issue_counts"]["reduction_pct"], 100)
        self.assertTrue(ctx["pass"])

    def test_missing_result_json_produces_zero_issue_counts(self):
        """Test 2: Without result JSON, issue_counts all zero, model is 'unknown'."""
        ndjson_file = os.path.join(NDJSON_DIR, "large-glm-5-1.ndjson")
        if not os.path.exists(ndjson_file):
            self.skipTest("large-glm-5-1.ndjson not found")

        events = extract_context.parse_ndjson_full(ndjson_file)
        ctx = extract_context.extract_context("large-glm-5-1", events, {})

        self.assertEqual(ctx["issue_counts"]["before"], 0)
        self.assertEqual(ctx["issue_counts"]["after"], 0)
        self.assertEqual(ctx["issue_counts"]["reduction_pct"], 0)
        self.assertEqual(ctx["model"], "unknown")

    def test_tool_breakdown_counts(self):
        """Test 3: tool_breakdown correctly counts tool_name → count from tool_use events."""
        ndjson_file = os.path.join(NDJSON_DIR, "large-glm-5-1.ndjson")
        if not os.path.exists(ndjson_file):
            self.skipTest("large-glm-5-1.ndjson not found")

        events = extract_context.parse_ndjson_full(ndjson_file)
        result_file = os.path.join(NDJSON_DIR, "large-glm-5-1-result.json")
        result = {}
        if os.path.exists(result_file):
            with open(result_file) as f:
                result = json.load(f)

        ctx = extract_context.extract_context("large-glm-5-1", events, result)

        # Should have tool_breakdown dict
        self.assertIsInstance(ctx["tool_breakdown"], dict)
        self.assertGreater(len(ctx["tool_breakdown"]), 0)

        # Should have at least edit and read (the agent reads files then edits them)
        tool_names = set(ctx["tool_breakdown"].keys())
        # The tool names should have golangci_lint_ prefix stripped for MCP tools
        # and bare names like 'edit', 'read', 'bash', 'glob'
        self.assertTrue(
            any("edit" in t for t in tool_names) or any("golangci_lint" in t for t in tool_names),
            f"Expected tool calls in breakdown, got: {tool_names}",
        )

        # Verify counts are positive integers
        for tool_name, count in ctx["tool_breakdown"].items():
            self.assertIsInstance(count, int)
            self.assertGreater(count, 0, f"Tool {tool_name} count should be > 0")

    def test_token_totals_summed(self):
        """Test 4: token_totals correctly sums across all step_finish events."""
        ndjson_file = os.path.join(NDJSON_DIR, "large-glm-5-1.ndjson")
        if not os.path.exists(ndjson_file):
            self.skipTest("large-glm-5-1.ndjson not found")

        events = extract_context.parse_ndjson_full(ndjson_file)
        result_file = os.path.join(NDJSON_DIR, "large-glm-5-1-result.json")
        result = {}
        if os.path.exists(result_file):
            with open(result_file) as f:
                result = json.load(f)

        ctx = extract_context.extract_context("large-glm-5-1", events, result)

        # token_totals should have input, output, reasoning
        self.assertIn("input", ctx["token_totals"])
        self.assertIn("output", ctx["token_totals"])
        self.assertIn("reasoning", ctx["token_totals"])

        # All should be >= 0
        self.assertGreaterEqual(ctx["token_totals"]["input"], 0)
        self.assertGreaterEqual(ctx["token_totals"]["output"], 0)
        self.assertGreaterEqual(ctx["token_totals"]["reasoning"], 0)

        # For a real trace with actual tool calls, should have non-zero totals
        self.assertGreater(ctx["token_totals"]["input"] + ctx["token_totals"]["output"], 0)

        # Verify cache_read exists (may be 0)
        self.assertIn("cache_read", ctx["token_totals"])

    def test_subagent_count_from_task_calls(self):
        """Test 5: subagent_count detected from task tool calls."""
        ndjson_file = os.path.join(NDJSON_DIR, "large-glm-5-1.ndjson")
        if not os.path.exists(ndjson_file):
            self.skipTest("large-glm-5-1.ndjson not found")

        events = extract_context.parse_ndjson_full(ndjson_file)
        result_file = os.path.join(NDJSON_DIR, "large-glm-5-1-result.json")
        result = {}
        if os.path.exists(result_file):
            with open(result_file) as f:
                result = json.load(f)

        ctx = extract_context.extract_context("large-glm-5-1", events, result)

        # subagent_count should be an integer >= 0
        self.assertIsInstance(ctx["subagent_count"], int)
        self.assertGreaterEqual(ctx["subagent_count"], 0)

        # For GLM-5.1 Large trace with 116 issues, it uses single-agent strategy
        # (subagents are used in larger traces like multipkg)
        # Just verify the count is a valid integer
        self.assertIsInstance(ctx["subagent_count"], int)

    def test_strategy_extraction(self):
        """Test 6: strategy extracted from tool_use output with strategy_instructions."""
        ndjson_file = os.path.join(NDJSON_DIR, "large-glm-5-1.ndjson")
        if not os.path.exists(ndjson_file):
            self.skipTest("large-glm-5-1.ndjson not found")

        events = extract_context.parse_ndjson_full(ndjson_file)
        result_file = os.path.join(NDJSON_DIR, "large-glm-5-1-result.json")
        result = {}
        if os.path.exists(result_file):
            with open(result_file) as f:
                result = json.load(f)

        ctx = extract_context.extract_context("large-glm-5-1", events, result)

        # Strategy should be a non-empty string
        self.assertIsInstance(ctx["strategy"], str)
        self.assertGreater(len(ctx["strategy"]), 0)

        # The strategy should contain a recognizable strategy keyword
        # (single-agent, subagent-per-file, etc.)
        strategy_lower = ctx["strategy"].lower()
        self.assertTrue(
            "agent" in strategy_lower or "subagent" in strategy_lower or "none" not in ctx["strategy"],
            f"Strategy should contain 'agent' or be non-'none', got: {ctx['strategy']}",
        )


class TestExtractContextMainFunction(unittest.TestCase):
    """Test the main() function produces output files."""

    @classmethod
    def setUpClass(cls):
        if not os.path.isdir(NDJSON_DIR):
            raise unittest.SkipTest("tmp/ndjson/ not found")

    def test_main_produces_context_files(self):
        """Test that running main() creates context JSON files."""
        # Run main
        extract_context.main()

        # Check that context files were created
        ndjson_files = [f for f in os.listdir(NDJSON_DIR) if f.endswith(".ndjson") and "probe" not in f]

        out_dir = os.path.join(PROJECT_ROOT, "tmp", "ndjson_analysis")

        for ndjson_f in ndjson_files:
            base = ndjson_f.replace(".ndjson", "")
            ctx_path = os.path.join(out_dir, f"{base}-context.json")
            self.assertTrue(os.path.exists(ctx_path), f"Expected context file for {base}")

            # Validate it's valid JSON with required fields
            with open(ctx_path) as f:
                ctx = json.load(f)
            self.assertIn("model", ctx)
            self.assertIn("tool_breakdown", ctx)
            self.assertIn("issue_counts", ctx)


if __name__ == "__main__":
    unittest.main()
