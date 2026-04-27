# Graph Report - guides  (2026-04-25)

## Corpus Check
- Large corpus: 630 files · ~92,124 words. Semantic extraction will be expensive (many Claude tokens). Consider running on a subfolder, or use --no-semantic to run AST-only.

## Summary
- 629 nodes · 2232 edges · 204 communities detected
- Extraction: 26% EXTRACTED · 74% INFERRED · 0% AMBIGUOUS · INFERRED: 1661 edges (avg confidence: 0.59)
- Token cost: 0 input · 0 output

## God Nodes (most connected - your core abstractions)
1. `G104` - 66 edges
2. `err113` - 51 edges
3. `G704` - 46 edges
4. `revive: unhandled-error` - 45 edges
5. `revive: deep-exit` - 43 edges
6. `exhaustive` - 42 edges
7. `cyclop` - 38 edges
8. `errcheck` - 37 edges
9. `gocritic: exitAfterDefer` - 37 edges
10. `staticcheck: SA4017` - 37 edges

## Surprising Connections (you probably didn't know these)
- `errchkjson` --semantically_similar_to--> `G104`  [INFERRED] [semantically similar]
  guides/errchkjson.md → guides/gosec/G104.md
- `arangolint` --semantically_similar_to--> `gocritic: sloppyTypeAssert`  [INFERRED] [semantically similar]
  guides/arangolint.md → guides/gocritic/sloppyTypeAssert.md
- `arangolint` --semantically_similar_to--> `G709`  [INFERRED] [semantically similar]
  guides/arangolint.md → guides/gosec/G709.md
- `arangolint` --semantically_similar_to--> `revive: unchecked-type-assertion`  [INFERRED] [semantically similar]
  guides/arangolint.md → guides/revive/unchecked-type-assertion.md
- `arangolint` --semantically_similar_to--> `staticcheck: ST1016`  [INFERRED] [semantically similar]
  guides/arangolint.md → guides/staticcheck/ST1016.md

## Hyperedges (group relationships)
- **Error Handling Linter Cluster** — errcheck, wrapcheck, err113, errname, errorlint_errorf, errorlint_asserts, errorlint_comparison, govet_printf, nilerr, nilnesserr [INFERRED 0.80]
- **Security Linter Cluster** — gosec_G101, gosec_G102, gosec_G103, gosec_G104, gosec_G201, gosec_G202, gosec_G203, gosec_G204, gosec_G301, gosec_G304, gosec_G401, gosec_G402, gosec_G501, gosec_G502, gosec_G505, gosec_G601, bidichk, durationcheck, noctx [INFERRED 0.80]
- **Complexity Linter Cluster** — cyclop, funlen, gocognit, gocyclo, maintidx, nestif, revive_cognitive-complexity, revive_cyclomatic, revive_max-control-nesting, revive_function-length, gocritic_nestingReduce, gocritic_ifElseChain [INFERRED 0.80]
- **Testing Linter Cluster** — tparallel, paralleltest, testpackage, testableexamples, thelper, testifylint_nil-compare, testifylint_expected-actual, testifylint_require-error, testifylint_error-nil, testifylint_bool-compare, testifylint_go-require, testifylint_error-as, testifylint_useless-assert, ginkgolinter_nil-assertion, ginkgolinter_error-assertion, ginkgolinter_compare-assertion, ginkgolinter_succeed-matcher [INFERRED 0.80]
- **Style Formatting Linter Cluster** — gofmt, gofumpt, decorder, nlreturn, wsl_v5, godot, goheader, misspell, lll, revive_file-length-limit, revive_line-length-limit, revive_comment-spacings, revive_imports-blocklist, gocritic_commentFormatting, gocritic_commentedOutCode [INFERRED 0.80]
- **Concurrency Linter Cluster** — bodyclose, contextcheck, makezero, govet_lostcancel, govet_atomic, govet_waitgroup, govet_tests, govet_testinggoroutine, revive_datarace, revive_waitgroup-by-value, containedctx, spancheck [INFERRED 0.80]
- **Performance Linter Cluster** — prealloc, wastedassign, ineffassign, unconvert, perfsprint, govet_copylocks, gocritic_rangeValCopy, gocritic_rangeExprCopy, gocritic_appendCombine, gocritic_sloppyReassign, fatcontext, gocritic_hugeParam, gocritic_indexAlloc [INFERRED 0.80]
- **Static Analysis Linter Cluster** — unused, unparam, govet_unreachable, govet_bools, govet_assign, govet_composites, govet_stdmethods, govet_errorsas, staticcheck_SA4000, staticcheck_SA4006, staticcheck_SA4010, staticcheck_SA4016, staticcheck_SA4019, gocritic_dupArg, gocritic_dupSubExpr, gocritic_dupBranchBody [INFERRED 0.80]

## Communities

### Community 0 - "Security Auditing"
Cohesion: 0.12
Nodes (66): bidichk, durationcheck, G101, G102, G103, G105, G106, G107 (+58 more)

### Community 1 - "Error Handling"
Cohesion: 0.24
Nodes (62): err113, errcheck, errname, errorlint: asserts, errorlint: comparison, errorlint: errorf, exhaustive, gochecksumtype (+54 more)

### Community 2 - "Concurrency Safety"
Cohesion: 0.17
Nodes (56): bodyclose, containedctx, contextcheck, gocritic: badLock, gocritic: badSyncOnceFunc, gocritic: deferInLoop, gocritic: exitAfterDefer, gocritic: syncMapLoadAndDelete (+48 more)

### Community 3 - "Style and Formatting"
Cohesion: 0.21
Nodes (50): asciicheck, decorder, dupword, forbidigo, funcorder, gocheckcompilerdirectives, gochecknoglobals, gochecknoinits (+42 more)

### Community 4 - "Dead Code Detection"
Cohesion: 0.21
Nodes (40): gocritic: argOrder, gocritic: caseOrder, gocritic: dupArg, gocritic: dupBranchBody, gocritic: dupCase, gocritic: dupOption, gocritic: dupSubExpr, gocritic: evalOrder (+32 more)

### Community 5 - "Code Complexity"
Cohesion: 0.21
Nodes (33): arangolint, cyclop, dogsled, dupl, errchkjson, funlen, gocognit, gocritic: defaultCaseOrder (+25 more)

### Community 6 - "Performance Optimization"
Cohesion: 0.25
Nodes (29): fatcontext, gocritic: appendAssign, gocritic: appendCombine, gocritic: badCond, gocritic: badSorting, gocritic: builtinShadow, gocritic: builtinShadowDecl, gocritic: hugeParam (+21 more)

### Community 7 - "Code Simplification"
Cohesion: 0.09
Nodes (27): gocritic: assignOp, gocritic: boolExprSimplify, gocritic: captLocal, gocritic: deferUnlambda, gocritic: emptyStringTest, gocritic: exposedSyncMutex, gocritic: filepathJoin, gocritic: hexLiteral (+19 more)

### Community 8 - "Deprecated API Patterns"
Cohesion: 0.1
Nodes (24): depguard, staticcheck: SA1000, staticcheck: SA1010, staticcheck: SA1019, staticcheck: SA4001, staticcheck: SA4003, staticcheck: SA4011, staticcheck: SA4012 (+16 more)

### Community 9 - "Testing Frameworks"
Cohesion: 0.92
Nodes (18): ginkgolinter: compare-assertion, ginkgolinter: error-assertion, ginkgolinter: nil-assertion, ginkgolinter: succeed-matcher, paralleltest, testableexamples, testifylint: bool-compare, testifylint: error-as (+10 more)

### Community 10 - "Community 10"
Cohesion: 0.47
Nodes (13): embeddedstructfieldcheck, exhaustruct, iface, inamedparam, musttag, recvcheck, staticcheck: SA4029, staticcheck: SA5007 (+5 more)

### Community 11 - "Community 11"
Cohesion: 0.47
Nodes (6): gocritic: badRegexp, gocritic: dynamicFmtString, gocritic: regexpMust, gocritic: regexpPattern, gocritic: regexpSimplify, gocritic: sprintfQuotedString

### Community 12 - "Community 12"
Cohesion: 0.47
Nodes (6): staticcheck: SA1021, staticcheck: SA1023, staticcheck: SA4020, staticcheck: SA4023, staticcheck: SA4031, staticcheck: SA5010

### Community 13 - "Community 13"
Cohesion: 0.83
Nodes (4): govet: asmdecl, govet: cgocall, govet: framepointer, govet: unsafeptr

### Community 14 - "Community 14"
Cohesion: 1.0
Nodes (3): govet: buildtag, govet: directive, govet: stdversion

### Community 15 - "Community 15"
Cohesion: 0.67
Nodes (3): staticcheck: SA1003, staticcheck: SA1011, staticcheck: SA1014

### Community 16 - "Community 16"
Cohesion: 1.0
Nodes (2): staticcheck: SA1026, staticcheck: SA1028

### Community 17 - "Community 17"
Cohesion: 1.0
Nodes (1): asasalint

### Community 18 - "Community 18"
Cohesion: 1.0
Nodes (1): canonicalheader

### Community 19 - "Community 19"
Cohesion: 1.0
Nodes (1): copyloopvar

### Community 20 - "Community 20"
Cohesion: 1.0
Nodes (1): exptostd

### Community 21 - "Community 21"
Cohesion: 1.0
Nodes (1): forcetypeassert

### Community 22 - "Community 22"
Cohesion: 1.0
Nodes (1): ginkgolinter: async-assertion

### Community 23 - "Community 23"
Cohesion: 1.0
Nodes (1): ginkgolinter: async-intervals

### Community 24 - "Community 24"
Cohesion: 1.0
Nodes (1): ginkgolinter: expect-to

### Community 25 - "Community 25"
Cohesion: 1.0
Nodes (1): ginkgolinter: focus-container

### Community 26 - "Community 26"
Cohesion: 1.0
Nodes (1): ginkgolinter: have-len-zero

### Community 27 - "Community 27"
Cohesion: 1.0
Nodes (1): ginkgolinter: len-assertion

### Community 28 - "Community 28"
Cohesion: 1.0
Nodes (1): ginkgolinter: spec-pollution

### Community 29 - "Community 29"
Cohesion: 1.0
Nodes (1): ginkgolinter: type-compare

### Community 30 - "Community 30"
Cohesion: 1.0
Nodes (1): gocritic: equalFold

### Community 31 - "Community 31"
Cohesion: 1.0
Nodes (1): gocritic: preferDecodeRune

### Community 32 - "Community 32"
Cohesion: 1.0
Nodes (1): gocritic: preferWriteByte

### Community 33 - "Community 33"
Cohesion: 1.0
Nodes (1): gocritic: ruleguard

### Community 34 - "Community 34"
Cohesion: 1.0
Nodes (1): gocritic: sliceClear

### Community 35 - "Community 35"
Cohesion: 1.0
Nodes (1): gocritic: stringXbytes

### Community 36 - "Community 36"
Cohesion: 1.0
Nodes (1): gocritic: zeroByteRepeat

### Community 37 - "Community 37"
Cohesion: 1.0
Nodes (1): grouper: const

### Community 38 - "Community 38"
Cohesion: 1.0
Nodes (1): grouper: import

### Community 39 - "Community 39"
Cohesion: 1.0
Nodes (1): grouper: type

### Community 40 - "Community 40"
Cohesion: 1.0
Nodes (1): grouper: var

### Community 41 - "Community 41"
Cohesion: 1.0
Nodes (1): intrange

### Community 42 - "Community 42"
Cohesion: 1.0
Nodes (1): iotamixing

### Community 43 - "Community 43"
Cohesion: 1.0
Nodes (1): mirror

### Community 44 - "Community 44"
Cohesion: 1.0
Nodes (1): modernize: loopvar

### Community 45 - "Community 45"
Cohesion: 1.0
Nodes (1): modernize: maprange

### Community 46 - "Community 46"
Cohesion: 1.0
Nodes (1): modernize: mapval

### Community 47 - "Community 47"
Cohesion: 1.0
Nodes (1): modernize: reloop

### Community 48 - "Community 48"
Cohesion: 1.0
Nodes (1): modernize: simplifyrange

### Community 49 - "Community 49"
Cohesion: 1.0
Nodes (1): modernize: sliceclear

### Community 50 - "Community 50"
Cohesion: 1.0
Nodes (1): modernize: slicesort

### Community 51 - "Community 51"
Cohesion: 1.0
Nodes (1): modernize: sortfunc

### Community 52 - "Community 52"
Cohesion: 1.0
Nodes (1): modernize: stringappend

### Community 53 - "Community 53"
Cohesion: 1.0
Nodes (1): revive: add-constant

### Community 54 - "Community 54"
Cohesion: 1.0
Nodes (1): revive: argument-limit

### Community 55 - "Community 55"
Cohesion: 1.0
Nodes (1): revive: banned-characters

### Community 56 - "Community 56"
Cohesion: 1.0
Nodes (1): revive: bare-return

### Community 57 - "Community 57"
Cohesion: 1.0
Nodes (1): revive: blank-imports

### Community 58 - "Community 58"
Cohesion: 1.0
Nodes (1): revive: bool-literal-in-expr

### Community 59 - "Community 59"
Cohesion: 1.0
Nodes (1): revive: comments-density

### Community 60 - "Community 60"
Cohesion: 1.0
Nodes (1): revive: confusing-naming

### Community 61 - "Community 61"
Cohesion: 1.0
Nodes (1): revive: confusing-results

### Community 62 - "Community 62"
Cohesion: 1.0
Nodes (1): revive: constant-logical-expr

### Community 63 - "Community 63"
Cohesion: 1.0
Nodes (1): revive: context-as-argument

### Community 64 - "Community 64"
Cohesion: 1.0
Nodes (1): revive: context-keys-type

### Community 65 - "Community 65"
Cohesion: 1.0
Nodes (1): revive: duplicated-imports

### Community 66 - "Community 66"
Cohesion: 1.0
Nodes (1): revive: early-return

### Community 67 - "Community 67"
Cohesion: 1.0
Nodes (1): revive: empty-block

### Community 68 - "Community 68"
Cohesion: 1.0
Nodes (1): revive: empty-lines

### Community 69 - "Community 69"
Cohesion: 1.0
Nodes (1): revive: enforce-map-style

### Community 70 - "Community 70"
Cohesion: 1.0
Nodes (1): revive: enforce-repeated-arg-type-style

### Community 71 - "Community 71"
Cohesion: 1.0
Nodes (1): revive: enforce-slice-style

### Community 72 - "Community 72"
Cohesion: 1.0
Nodes (1): revive: enforce-switch-style

### Community 73 - "Community 73"
Cohesion: 1.0
Nodes (1): revive: error-naming

### Community 74 - "Community 74"
Cohesion: 1.0
Nodes (1): revive: error-return

### Community 75 - "Community 75"
Cohesion: 1.0
Nodes (1): revive: exported

### Community 76 - "Community 76"
Cohesion: 1.0
Nodes (1): revive: file-header

### Community 77 - "Community 77"
Cohesion: 1.0
Nodes (1): revive: filename-format

### Community 78 - "Community 78"
Cohesion: 1.0
Nodes (1): revive: flag-parameter

### Community 79 - "Community 79"
Cohesion: 1.0
Nodes (1): revive: function-result-limit

### Community 80 - "Community 80"
Cohesion: 1.0
Nodes (1): revive: get-return

### Community 81 - "Community 81"
Cohesion: 1.0
Nodes (1): revive: identical-branches

### Community 82 - "Community 82"
Cohesion: 1.0
Nodes (1): revive: identical-ifelseif-branches

### Community 83 - "Community 83"
Cohesion: 1.0
Nodes (1): revive: identical-ifelseif-conditions

### Community 84 - "Community 84"
Cohesion: 1.0
Nodes (1): revive: identical-switch-branches

### Community 85 - "Community 85"
Cohesion: 1.0
Nodes (1): revive: identical-switch-conditions

### Community 86 - "Community 86"
Cohesion: 1.0
Nodes (1): revive: if-return

### Community 87 - "Community 87"
Cohesion: 1.0
Nodes (1): revive: import-alias-naming

### Community 88 - "Community 88"
Cohesion: 1.0
Nodes (1): revive: import-shadowing

### Community 89 - "Community 89"
Cohesion: 1.0
Nodes (1): revive: increment-decrement

### Community 90 - "Community 90"
Cohesion: 1.0
Nodes (1): revive: indent-error-flow

### Community 91 - "Community 91"
Cohesion: 1.0
Nodes (1): revive: inefficient-map-lookup

### Community 92 - "Community 92"
Cohesion: 1.0
Nodes (1): revive: max-public-structs

### Community 93 - "Community 93"
Cohesion: 1.0
Nodes (1): revive: modifies-parameter

### Community 94 - "Community 94"
Cohesion: 1.0
Nodes (1): revive: modifies-value-receiver

### Community 95 - "Community 95"
Cohesion: 1.0
Nodes (1): revive: nested-structs

### Community 96 - "Community 96"
Cohesion: 1.0
Nodes (1): revive: optimize-operands-order

### Community 97 - "Community 97"
Cohesion: 1.0
Nodes (1): revive: package-comments

### Community 98 - "Community 98"
Cohesion: 1.0
Nodes (1): revive: package-directory-mismatch

### Community 99 - "Community 99"
Cohesion: 1.0
Nodes (1): revive: package-naming

### Community 100 - "Community 100"
Cohesion: 1.0
Nodes (1): revive: range-val-address

### Community 101 - "Community 101"
Cohesion: 1.0
Nodes (1): revive: range

### Community 102 - "Community 102"
Cohesion: 1.0
Nodes (1): revive: receiver-naming

### Community 103 - "Community 103"
Cohesion: 1.0
Nodes (1): revive: redefines-builtin-id

### Community 104 - "Community 104"
Cohesion: 1.0
Nodes (1): revive: redundant-build-tag

### Community 105 - "Community 105"
Cohesion: 1.0
Nodes (1): revive: redundant-import-alias

### Community 106 - "Community 106"
Cohesion: 1.0
Nodes (1): revive: redundant-test-main-exit

### Community 107 - "Community 107"
Cohesion: 1.0
Nodes (1): revive: string-of-int

### Community 108 - "Community 108"
Cohesion: 1.0
Nodes (1): revive: struct-tag

### Community 109 - "Community 109"
Cohesion: 1.0
Nodes (1): revive: superfluous-else

### Community 110 - "Community 110"
Cohesion: 1.0
Nodes (1): revive: time-date

### Community 111 - "Community 111"
Cohesion: 1.0
Nodes (1): revive: time-naming

### Community 112 - "Community 112"
Cohesion: 1.0
Nodes (1): revive: unconditional-recursion

### Community 113 - "Community 113"
Cohesion: 1.0
Nodes (1): revive: unexported-naming

### Community 114 - "Community 114"
Cohesion: 1.0
Nodes (1): revive: unexported-return

### Community 115 - "Community 115"
Cohesion: 1.0
Nodes (1): revive: unnecessary-format

### Community 116 - "Community 116"
Cohesion: 1.0
Nodes (1): revive: unnecessary-if

### Community 117 - "Community 117"
Cohesion: 1.0
Nodes (1): revive: unnecessary-stmt

### Community 118 - "Community 118"
Cohesion: 1.0
Nodes (1): revive: unsecure-url-scheme

### Community 119 - "Community 119"
Cohesion: 1.0
Nodes (1): revive: unused-parameter

### Community 120 - "Community 120"
Cohesion: 1.0
Nodes (1): revive: unused-receiver

### Community 121 - "Community 121"
Cohesion: 1.0
Nodes (1): revive: use-slices-sort

### Community 122 - "Community 122"
Cohesion: 1.0
Nodes (1): revive: useless-break

### Community 123 - "Community 123"
Cohesion: 1.0
Nodes (1): revive: useless-fallthrough

### Community 124 - "Community 124"
Cohesion: 1.0
Nodes (1): revive: var-declaration

### Community 125 - "Community 125"
Cohesion: 1.0
Nodes (1): revive: var-naming

### Community 126 - "Community 126"
Cohesion: 1.0
Nodes (1): staticcheck: QF1001

### Community 127 - "Community 127"
Cohesion: 1.0
Nodes (1): staticcheck: QF1002

### Community 128 - "Community 128"
Cohesion: 1.0
Nodes (1): staticcheck: QF1003

### Community 129 - "Community 129"
Cohesion: 1.0
Nodes (1): staticcheck: QF1006

### Community 130 - "Community 130"
Cohesion: 1.0
Nodes (1): staticcheck: QF1007

### Community 131 - "Community 131"
Cohesion: 1.0
Nodes (1): staticcheck: QF1008

### Community 132 - "Community 132"
Cohesion: 1.0
Nodes (1): staticcheck: QF1009

### Community 133 - "Community 133"
Cohesion: 1.0
Nodes (1): staticcheck: QF1010

### Community 134 - "Community 134"
Cohesion: 1.0
Nodes (1): staticcheck: QF1011

### Community 135 - "Community 135"
Cohesion: 1.0
Nodes (1): staticcheck: QF1012

### Community 136 - "Community 136"
Cohesion: 1.0
Nodes (1): staticcheck: S1000

### Community 137 - "Community 137"
Cohesion: 1.0
Nodes (1): staticcheck: S1001

### Community 138 - "Community 138"
Cohesion: 1.0
Nodes (1): staticcheck: S1002

### Community 139 - "Community 139"
Cohesion: 1.0
Nodes (1): staticcheck: S1003

### Community 140 - "Community 140"
Cohesion: 1.0
Nodes (1): staticcheck: S1004

### Community 141 - "Community 141"
Cohesion: 1.0
Nodes (1): staticcheck: S1006

### Community 142 - "Community 142"
Cohesion: 1.0
Nodes (1): staticcheck: S1007

### Community 143 - "Community 143"
Cohesion: 1.0
Nodes (1): staticcheck: S1008

### Community 144 - "Community 144"
Cohesion: 1.0
Nodes (1): staticcheck: S1009

### Community 145 - "Community 145"
Cohesion: 1.0
Nodes (1): staticcheck: S1010

### Community 146 - "Community 146"
Cohesion: 1.0
Nodes (1): staticcheck: S1011

### Community 147 - "Community 147"
Cohesion: 1.0
Nodes (1): staticcheck: S1012

### Community 148 - "Community 148"
Cohesion: 1.0
Nodes (1): staticcheck: S1013

### Community 149 - "Community 149"
Cohesion: 1.0
Nodes (1): staticcheck: S1014

### Community 150 - "Community 150"
Cohesion: 1.0
Nodes (1): staticcheck: S1015

### Community 151 - "Community 151"
Cohesion: 1.0
Nodes (1): staticcheck: S1016

### Community 152 - "Community 152"
Cohesion: 1.0
Nodes (1): staticcheck: S1017

### Community 153 - "Community 153"
Cohesion: 1.0
Nodes (1): staticcheck: S1018

### Community 154 - "Community 154"
Cohesion: 1.0
Nodes (1): staticcheck: S1019

### Community 155 - "Community 155"
Cohesion: 1.0
Nodes (1): staticcheck: S1020

### Community 156 - "Community 156"
Cohesion: 1.0
Nodes (1): staticcheck: S1021

### Community 157 - "Community 157"
Cohesion: 1.0
Nodes (1): staticcheck: S1023

### Community 158 - "Community 158"
Cohesion: 1.0
Nodes (1): staticcheck: S1024

### Community 159 - "Community 159"
Cohesion: 1.0
Nodes (1): staticcheck: S1025

### Community 160 - "Community 160"
Cohesion: 1.0
Nodes (1): staticcheck: S1026

### Community 161 - "Community 161"
Cohesion: 1.0
Nodes (1): staticcheck: S1027

### Community 162 - "Community 162"
Cohesion: 1.0
Nodes (1): staticcheck: S1030

### Community 163 - "Community 163"
Cohesion: 1.0
Nodes (1): staticcheck: S1031

### Community 164 - "Community 164"
Cohesion: 1.0
Nodes (1): staticcheck: S1032

### Community 165 - "Community 165"
Cohesion: 1.0
Nodes (1): staticcheck: S1033

### Community 166 - "Community 166"
Cohesion: 1.0
Nodes (1): staticcheck: S1034

### Community 167 - "Community 167"
Cohesion: 1.0
Nodes (1): staticcheck: S1035

### Community 168 - "Community 168"
Cohesion: 1.0
Nodes (1): staticcheck: S1036

### Community 169 - "Community 169"
Cohesion: 1.0
Nodes (1): staticcheck: S1037

### Community 170 - "Community 170"
Cohesion: 1.0
Nodes (1): staticcheck: S1038

### Community 171 - "Community 171"
Cohesion: 1.0
Nodes (1): staticcheck: S1039

### Community 172 - "Community 172"
Cohesion: 1.0
Nodes (1): staticcheck: ST1000

### Community 173 - "Community 173"
Cohesion: 1.0
Nodes (1): staticcheck: ST1001

### Community 174 - "Community 174"
Cohesion: 1.0
Nodes (1): staticcheck: ST1002

### Community 175 - "Community 175"
Cohesion: 1.0
Nodes (1): staticcheck: ST1003

### Community 176 - "Community 176"
Cohesion: 1.0
Nodes (1): staticcheck: ST1005

### Community 177 - "Community 177"
Cohesion: 1.0
Nodes (1): staticcheck: ST1006

### Community 178 - "Community 178"
Cohesion: 1.0
Nodes (1): staticcheck: ST1007

### Community 179 - "Community 179"
Cohesion: 1.0
Nodes (1): staticcheck: ST1009

### Community 180 - "Community 180"
Cohesion: 1.0
Nodes (1): staticcheck: ST1010

### Community 181 - "Community 181"
Cohesion: 1.0
Nodes (1): staticcheck: ST1011

### Community 182 - "Community 182"
Cohesion: 1.0
Nodes (1): staticcheck: ST1012

### Community 183 - "Community 183"
Cohesion: 1.0
Nodes (1): staticcheck: ST1013

### Community 184 - "Community 184"
Cohesion: 1.0
Nodes (1): staticcheck: ST1014

### Community 185 - "Community 185"
Cohesion: 1.0
Nodes (1): staticcheck: ST1015

### Community 186 - "Community 186"
Cohesion: 1.0
Nodes (1): staticcheck: ST1017

### Community 187 - "Community 187"
Cohesion: 1.0
Nodes (1): staticcheck: ST1018

### Community 188 - "Community 188"
Cohesion: 1.0
Nodes (1): staticcheck: ST1019

### Community 189 - "Community 189"
Cohesion: 1.0
Nodes (1): staticcheck: ST1020

### Community 190 - "Community 190"
Cohesion: 1.0
Nodes (1): staticcheck: ST1021

### Community 191 - "Community 191"
Cohesion: 1.0
Nodes (1): staticcheck: ST1023

### Community 192 - "Community 192"
Cohesion: 1.0
Nodes (1): testifylint: blank-import

### Community 193 - "Community 193"
Cohesion: 1.0
Nodes (1): testifylint: compares

### Community 194 - "Community 194"
Cohesion: 1.0
Nodes (1): testifylint: contains-unnecessary-format

### Community 195 - "Community 195"
Cohesion: 1.0
Nodes (1): testifylint: empty

### Community 196 - "Community 196"
Cohesion: 1.0
Nodes (1): testifylint: float-compare

### Community 197 - "Community 197"
Cohesion: 1.0
Nodes (1): testifylint: formatter

### Community 198 - "Community 198"
Cohesion: 1.0
Nodes (1): testifylint: len

### Community 199 - "Community 199"
Cohesion: 1.0
Nodes (1): testifylint: suite-broken-parallel

### Community 200 - "Community 200"
Cohesion: 1.0
Nodes (1): testifylint: suite-dont-use-pkg

### Community 201 - "Community 201"
Cohesion: 1.0
Nodes (1): testifylint: suite-extra-assert-call

### Community 202 - "Community 202"
Cohesion: 1.0
Nodes (1): testifylint: suite-method-signature

### Community 203 - "Community 203"
Cohesion: 1.0
Nodes (1): testifylint: suite-thelper

## Knowledge Gaps
- **209 isolated node(s):** `asasalint`, `canonicalheader`, `copyloopvar`, `depguard`, `exptostd` (+204 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Community 16`** (2 nodes): `staticcheck: SA1026`, `staticcheck: SA1028`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 17`** (1 nodes): `asasalint`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 18`** (1 nodes): `canonicalheader`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 19`** (1 nodes): `copyloopvar`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 20`** (1 nodes): `exptostd`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 21`** (1 nodes): `forcetypeassert`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 22`** (1 nodes): `ginkgolinter: async-assertion`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 23`** (1 nodes): `ginkgolinter: async-intervals`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 24`** (1 nodes): `ginkgolinter: expect-to`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 25`** (1 nodes): `ginkgolinter: focus-container`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 26`** (1 nodes): `ginkgolinter: have-len-zero`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 27`** (1 nodes): `ginkgolinter: len-assertion`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 28`** (1 nodes): `ginkgolinter: spec-pollution`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 29`** (1 nodes): `ginkgolinter: type-compare`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 30`** (1 nodes): `gocritic: equalFold`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 31`** (1 nodes): `gocritic: preferDecodeRune`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 32`** (1 nodes): `gocritic: preferWriteByte`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 33`** (1 nodes): `gocritic: ruleguard`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 34`** (1 nodes): `gocritic: sliceClear`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 35`** (1 nodes): `gocritic: stringXbytes`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 36`** (1 nodes): `gocritic: zeroByteRepeat`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 37`** (1 nodes): `grouper: const`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 38`** (1 nodes): `grouper: import`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 39`** (1 nodes): `grouper: type`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 40`** (1 nodes): `grouper: var`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 41`** (1 nodes): `intrange`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 42`** (1 nodes): `iotamixing`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 43`** (1 nodes): `mirror`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 44`** (1 nodes): `modernize: loopvar`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 45`** (1 nodes): `modernize: maprange`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 46`** (1 nodes): `modernize: mapval`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 47`** (1 nodes): `modernize: reloop`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 48`** (1 nodes): `modernize: simplifyrange`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 49`** (1 nodes): `modernize: sliceclear`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 50`** (1 nodes): `modernize: slicesort`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 51`** (1 nodes): `modernize: sortfunc`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 52`** (1 nodes): `modernize: stringappend`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 53`** (1 nodes): `revive: add-constant`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 54`** (1 nodes): `revive: argument-limit`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 55`** (1 nodes): `revive: banned-characters`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 56`** (1 nodes): `revive: bare-return`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 57`** (1 nodes): `revive: blank-imports`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 58`** (1 nodes): `revive: bool-literal-in-expr`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 59`** (1 nodes): `revive: comments-density`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 60`** (1 nodes): `revive: confusing-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 61`** (1 nodes): `revive: confusing-results`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 62`** (1 nodes): `revive: constant-logical-expr`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 63`** (1 nodes): `revive: context-as-argument`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 64`** (1 nodes): `revive: context-keys-type`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 65`** (1 nodes): `revive: duplicated-imports`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 66`** (1 nodes): `revive: early-return`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 67`** (1 nodes): `revive: empty-block`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 68`** (1 nodes): `revive: empty-lines`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 69`** (1 nodes): `revive: enforce-map-style`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 70`** (1 nodes): `revive: enforce-repeated-arg-type-style`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 71`** (1 nodes): `revive: enforce-slice-style`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 72`** (1 nodes): `revive: enforce-switch-style`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 73`** (1 nodes): `revive: error-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 74`** (1 nodes): `revive: error-return`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 75`** (1 nodes): `revive: exported`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 76`** (1 nodes): `revive: file-header`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 77`** (1 nodes): `revive: filename-format`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 78`** (1 nodes): `revive: flag-parameter`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 79`** (1 nodes): `revive: function-result-limit`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 80`** (1 nodes): `revive: get-return`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 81`** (1 nodes): `revive: identical-branches`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 82`** (1 nodes): `revive: identical-ifelseif-branches`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 83`** (1 nodes): `revive: identical-ifelseif-conditions`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 84`** (1 nodes): `revive: identical-switch-branches`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 85`** (1 nodes): `revive: identical-switch-conditions`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 86`** (1 nodes): `revive: if-return`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 87`** (1 nodes): `revive: import-alias-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 88`** (1 nodes): `revive: import-shadowing`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 89`** (1 nodes): `revive: increment-decrement`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 90`** (1 nodes): `revive: indent-error-flow`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 91`** (1 nodes): `revive: inefficient-map-lookup`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 92`** (1 nodes): `revive: max-public-structs`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 93`** (1 nodes): `revive: modifies-parameter`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 94`** (1 nodes): `revive: modifies-value-receiver`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 95`** (1 nodes): `revive: nested-structs`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 96`** (1 nodes): `revive: optimize-operands-order`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 97`** (1 nodes): `revive: package-comments`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 98`** (1 nodes): `revive: package-directory-mismatch`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 99`** (1 nodes): `revive: package-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 100`** (1 nodes): `revive: range-val-address`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 101`** (1 nodes): `revive: range`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 102`** (1 nodes): `revive: receiver-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 103`** (1 nodes): `revive: redefines-builtin-id`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 104`** (1 nodes): `revive: redundant-build-tag`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 105`** (1 nodes): `revive: redundant-import-alias`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 106`** (1 nodes): `revive: redundant-test-main-exit`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 107`** (1 nodes): `revive: string-of-int`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 108`** (1 nodes): `revive: struct-tag`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 109`** (1 nodes): `revive: superfluous-else`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 110`** (1 nodes): `revive: time-date`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 111`** (1 nodes): `revive: time-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 112`** (1 nodes): `revive: unconditional-recursion`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 113`** (1 nodes): `revive: unexported-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 114`** (1 nodes): `revive: unexported-return`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 115`** (1 nodes): `revive: unnecessary-format`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 116`** (1 nodes): `revive: unnecessary-if`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 117`** (1 nodes): `revive: unnecessary-stmt`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 118`** (1 nodes): `revive: unsecure-url-scheme`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 119`** (1 nodes): `revive: unused-parameter`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 120`** (1 nodes): `revive: unused-receiver`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 121`** (1 nodes): `revive: use-slices-sort`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 122`** (1 nodes): `revive: useless-break`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 123`** (1 nodes): `revive: useless-fallthrough`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 124`** (1 nodes): `revive: var-declaration`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 125`** (1 nodes): `revive: var-naming`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 126`** (1 nodes): `staticcheck: QF1001`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 127`** (1 nodes): `staticcheck: QF1002`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 128`** (1 nodes): `staticcheck: QF1003`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 129`** (1 nodes): `staticcheck: QF1006`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 130`** (1 nodes): `staticcheck: QF1007`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 131`** (1 nodes): `staticcheck: QF1008`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 132`** (1 nodes): `staticcheck: QF1009`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 133`** (1 nodes): `staticcheck: QF1010`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 134`** (1 nodes): `staticcheck: QF1011`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 135`** (1 nodes): `staticcheck: QF1012`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 136`** (1 nodes): `staticcheck: S1000`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 137`** (1 nodes): `staticcheck: S1001`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 138`** (1 nodes): `staticcheck: S1002`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 139`** (1 nodes): `staticcheck: S1003`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 140`** (1 nodes): `staticcheck: S1004`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 141`** (1 nodes): `staticcheck: S1006`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 142`** (1 nodes): `staticcheck: S1007`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 143`** (1 nodes): `staticcheck: S1008`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 144`** (1 nodes): `staticcheck: S1009`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 145`** (1 nodes): `staticcheck: S1010`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 146`** (1 nodes): `staticcheck: S1011`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 147`** (1 nodes): `staticcheck: S1012`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 148`** (1 nodes): `staticcheck: S1013`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 149`** (1 nodes): `staticcheck: S1014`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 150`** (1 nodes): `staticcheck: S1015`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 151`** (1 nodes): `staticcheck: S1016`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 152`** (1 nodes): `staticcheck: S1017`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 153`** (1 nodes): `staticcheck: S1018`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 154`** (1 nodes): `staticcheck: S1019`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 155`** (1 nodes): `staticcheck: S1020`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 156`** (1 nodes): `staticcheck: S1021`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 157`** (1 nodes): `staticcheck: S1023`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 158`** (1 nodes): `staticcheck: S1024`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 159`** (1 nodes): `staticcheck: S1025`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 160`** (1 nodes): `staticcheck: S1026`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 161`** (1 nodes): `staticcheck: S1027`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 162`** (1 nodes): `staticcheck: S1030`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 163`** (1 nodes): `staticcheck: S1031`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 164`** (1 nodes): `staticcheck: S1032`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 165`** (1 nodes): `staticcheck: S1033`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 166`** (1 nodes): `staticcheck: S1034`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 167`** (1 nodes): `staticcheck: S1035`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 168`** (1 nodes): `staticcheck: S1036`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 169`** (1 nodes): `staticcheck: S1037`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 170`** (1 nodes): `staticcheck: S1038`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 171`** (1 nodes): `staticcheck: S1039`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 172`** (1 nodes): `staticcheck: ST1000`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 173`** (1 nodes): `staticcheck: ST1001`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 174`** (1 nodes): `staticcheck: ST1002`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 175`** (1 nodes): `staticcheck: ST1003`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 176`** (1 nodes): `staticcheck: ST1005`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 177`** (1 nodes): `staticcheck: ST1006`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 178`** (1 nodes): `staticcheck: ST1007`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 179`** (1 nodes): `staticcheck: ST1009`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 180`** (1 nodes): `staticcheck: ST1010`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 181`** (1 nodes): `staticcheck: ST1011`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 182`** (1 nodes): `staticcheck: ST1012`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 183`** (1 nodes): `staticcheck: ST1013`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 184`** (1 nodes): `staticcheck: ST1014`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 185`** (1 nodes): `staticcheck: ST1015`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 186`** (1 nodes): `staticcheck: ST1017`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 187`** (1 nodes): `staticcheck: ST1018`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 188`** (1 nodes): `staticcheck: ST1019`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 189`** (1 nodes): `staticcheck: ST1020`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 190`** (1 nodes): `staticcheck: ST1021`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 191`** (1 nodes): `staticcheck: ST1023`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 192`** (1 nodes): `testifylint: blank-import`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 193`** (1 nodes): `testifylint: compares`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 194`** (1 nodes): `testifylint: contains-unnecessary-format`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 195`** (1 nodes): `testifylint: empty`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 196`** (1 nodes): `testifylint: float-compare`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 197`** (1 nodes): `testifylint: formatter`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 198`** (1 nodes): `testifylint: len`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 199`** (1 nodes): `testifylint: suite-broken-parallel`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 200`** (1 nodes): `testifylint: suite-dont-use-pkg`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 201`** (1 nodes): `testifylint: suite-extra-assert-call`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 202`** (1 nodes): `testifylint: suite-method-signature`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 203`** (1 nodes): `testifylint: suite-thelper`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `G104` connect `Error Handling` to `Security Auditing`, `Concurrency Safety`, `Style and Formatting`, `Dead Code Detection`, `Code Complexity`?**
  _High betweenness centrality (0.069) - this node is a cross-community bridge._
- **Why does `gocritic: exitAfterDefer` connect `Concurrency Safety` to `Error Handling`, `Style and Formatting`, `Dead Code Detection`, `Performance Optimization`, `Code Simplification`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Why does `lll` connect `Style and Formatting` to `Security Auditing`, `Error Handling`, `Code Complexity`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._
- **Are the 62 inferred relationships involving `G104` (e.g. with `G101` and `G102`) actually correct?**
  _`G104` has 62 INFERRED edges - model-reasoned connections that need verification._
- **Are the 50 inferred relationships involving `err113` (e.g. with `wrapcheck` and `errname`) actually correct?**
  _`err113` has 50 INFERRED edges - model-reasoned connections that need verification._
- **Are the 44 inferred relationships involving `G704` (e.g. with `bodyclose` and `gocritic: returnAfterHttpError`) actually correct?**
  _`G704` has 44 INFERRED edges - model-reasoned connections that need verification._
- **Are the 45 inferred relationships involving `revive: unhandled-error` (e.g. with `errcheck` and `gocritic: deferInLoop`) actually correct?**
  _`revive: unhandled-error` has 45 INFERRED edges - model-reasoned connections that need verification._