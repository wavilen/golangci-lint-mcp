# Cluster Analysis: Linter Relationship Research

Generated: 2026-04-25
Source: graphify-out/graph.json + GRAPH_REPORT.md
Purpose: Actionable input for Phase 57 related-tag curation

## Executive Summary

Graphify extracted 629 nodes and 2,232 edges from all guide files, revealing 10 major communities (≥13 members) and 8 category hyperedges. The top 10 communities cover 405/629 linter guides with cohesive thematic groupings: Security Auditing (66), Error Handling (62), Concurrency Safety (56), Style and Formatting (50), Dead Code Detection (40), Code Complexity (33), Performance Optimization (29), Code Simplification (27), Deprecated API Patterns (24), and Testing Frameworks (18). The remaining 224 guides form small clusters or singletons — these are the primary curation targets for Phase 57.

Key findings: 358/629 guides already have `<related>` tags, but many are trivially connected (only cross-referencing same-linter rules). The graph reveals 1,661 INFERRED relationships that are not captured in any `<related>` tag, representing significant curation opportunity.

## Cluster: Security Auditing

**Cohesion:** 0.12
**Members:** 66 linters

### Member List

- `bidichk` — detects dangerous unicode bidirectional characters
- `durationcheck` — detects multiplying time.Duration by time.Duration
- `noctx` — detects HTTP requests without context.Context
- `gosec/G101` — hardcoded credentials
- `gosec/G102` — network binding to all interfaces
- `gosec/G103` — unsafe block usage
- `gosec/G104` — audit errors not checked
- `gosec/G105` — math/big integer overflow
- `gosec/G106` — insecure SSH InsecureIgnoreHostKey
- `gosec/G107` — HTTP SSRF via URL from variable
- `gosec/G108` — profiling endpoint exposed
- `gosec/G109` — integer overflow in strconv.Atoi
- `gosec/G110` — decompression bomb potential
- `gosec/G111` — directory traversal via http.Dir
- `gosec/G112` — poor randomness source
- `gosec/G113` — unquoted SQL injection in CONCAT
- `gosec/G114` — improper cookie timeout
- `gosec/G115` — integer overflow in math/big
- `gosec/G116` — SQL injection via Append
- `gosec/G117` — weak cryptographic key size
- `gosec/G119` — weak pseudorandom number generator
- `gosec/G120` — TLS certificate verification disabled
- `gosec/G121` — hardcoded IP addresses
- `gosec/G123` — weak cipher suites
- `gosec/G124` — unrestricted file upload
- `gosec/G201` — SQL query construction via string concat
- `gosec/G202` — SQL query via string concatenation
- `gosec/G203` — HTML template injection
- `gosec/G204` — command execution from variable
- `gosec/G301` — poor file permissions (mkdir)
- `gosec/G302` — poor file permissions (chmod)
- `gosec/G303` — predictable random source (math/rand)
- `gosec/G304` — file path injection from variable
- `gosec/G305` — zip slip vulnerability
- `gosec/G306` — insecure file permissions (write)
- `gosec/G307` — insecure file permissions (create)
- `gosec/G401` — weak cryptographic primitive (MD5/SHA1)
- `gosec/G402` — insecure TLS configuration
- `gosec/G403` — weak AES key size
- `gosec/G404` — insecure random number source
- `gosec/G405` — weak cipher (DES/RC4/Blowfish)
- `gosec/G406` — weak key derivation (Simple DK)
- `gosec/G407` — potential directory traversal
- `gosec/G408` — weak SSH key exchange algorithm
- `gosec/G501` — blocked import (crypto/md5)
- `gosec/G502` — blocked import (crypto/des)
- `gosec/G503` — blocked import (crypto/rc4)
- `gosec/G504` — blocked import (net/http/cgi)
- `gosec/G505` — blocked import (crypto/sha1)
- `gosec/G506` — weak cookie without HttpOnly
- `gosec/G507` — weak cookie without Secure flag
- `gosec/G601` — implicit memory aliasing
- `gosec/G602` — slice memory aliasing via range
- `gosec/G701` — memory address manipulation (printf %p)
- `gosec/G702` — command injection via fmt.Sprintf
- `gosec/G703` — path injection via environment variable
- `gosec/G704` — improper input validation
- `gosec/G705` — CSS injection via untrusted input
- `gosec/G706` — potential DoS via deep nesting
- `gosec/G707` — unvalidated redirect
- `gosec/G708` — server-side request forgery via HTTP
- `gosec/G709` — potential integer overflow
- `revive/dot-imports` — (related: style overlap with security imports)
- `revive/string-format` — (related: format string patterns)
- `staticcheck/SA1007` — time.Sleep duration argument
- `staticcheck/SA1018` — invalid argument to Rand.Read
- `staticcheck/SA1020` — signal.Notify channel
- `staticcheck/SA1024` — signal.Notify duplicates
- `staticcheck/SA1031` — invalid argument to strconv

### Missing Relationships

- `contextcheck` is in the Concurrency cluster but is semantically related to `noctx` (both detect missing context.Context usage). Consider cross-referencing.
- `rowserrcheck` and `sqlclosecheck` are in Error Handling but share database security concerns with G104 (errors not checked). Consider referencing from `gosec/G104`.
- `errcheck` is a god node (37 edges) in Error Handling but directly relevant to `gosec/G104` (audit errors not checked). The existing `<related>` in `gosec/G104` already references `errcheck` — confirm this is intentional.
- `revive/unsecure-url-scheme` is in a different cluster but should be in this security cluster.

### Surprising Connections

- `errchkjson` ↔ `gosec/G104`: Both deal with unchecked error returns — errchkjson for JSON encoding specifically, G104 for audit trail. The graph correctly infers this latent coupling.
- `arangolint` ↔ `gosec/G709`: Both validate input patterns (arangolint for ArangoDB query safety, G709 for integer overflow). Semantic overlap in "input validation" domain.

### Curation Guidance

- `guides/gosec/G104.md` `<related>` currently lists `G109, errcheck` — appropriate. Keep.
- Add `noctx` to `<related>` in `guides/gosec/G107.md` (SSRF) — both deal with untrusted HTTP input paths.
- Add `gosec/G204` to `<related>` in `guides/gosec/G107.md` — command injection and SSRF are frequently co-exploited.
- `guides/gosec/G301.md`–`guides/gosec/G307.md` already cross-reference each other well (file permissions group). Keep as-is.
- Add `bidichk` to `<related>` in `guides/gosec/G101.md` — both detect hidden/trusted content issues.
- `guides/durationcheck.md` `<related>` is empty. Add `gosec/G109` (integer overflow) — both involve numeric overflow patterns.
- `guides/noctx.md` `<related>` currently lists `contextcheck, bodyclose, errcheck` — appropriate. Keep.
- The 61 gosec rules already have good internal cross-references. No major additions needed within gosec.

---

## Cluster: Error Handling

**Cohesion:** 0.24
**Members:** 62 linters

### Member List

- `err113` — errors.New comparison (error type checking)
- `errcheck` — unchecked error returns
- `errname` — error naming convention
- `errorlint/asserts` — error assertion patterns
- `errorlint/comparison` — error comparison patterns
- `errorlint/errorf` — error formatting with %w
- `exhaustive` — exhaustive enum/switch checking
- `gochecksumtype` — exhaustive interface type checking
- `gocritic/badCall` — suspicious function calls
- `gocritic/externalErrorReassign` — reassigning errors from external packages
- `gocritic/flagDeref` — nil pointer dereference on flag
- `gocritic/flagName` — flag naming convention
- `gocritic/importShadow` — import shadowing
- `gocritic/nilValReturn` — returning nil value after error check
- `gocritic/preferFprint` — prefer fmt.Fprint over fmt.Println
- `gocritic/preferStringWriter` — prefer StringWriter interface
- `gocritic/returnAfterHttpError` — missing return after HTTP error
- `gocritic/sloppyTypeAssert` — sloppy type assertion
- `gocritic/sqlQuery` — SQL query with unchecked error
- `gocritic/uncheckedInlineErr` — unchecked inline error
- `gocritic/weakCond` — weak condition
- `goprintffuncname` — printf-like function naming
- `gosec/G104` — audit errors not checked (cross-cluster bridge)
- `gosec/G704` — improper input validation (cross-cluster bridge)
- `govet/httpresponse` — unclosed HTTP response body
- `govet/ifaceassert` — impossible interface assertion
- `govet/printf` — printf format string issues
- `govet/unusedresult` — unused function result
- `modernize/errorf` — modernize error formatting
- `nilerr` — nil error return after error check
- `nilnesserr` — nilness follows error check
- `nilnil` — returning nil, nil
- `noinlineerr` — inline error checks
- `protogetter` — protobuf getter usage
- `revive/deep-exit` — deep exit calls (os.Exit, log.Fatal)
- `revive/error-strings` — error string convention
- `revive/errorf` — use fmt.Errorf for errors
- `revive/unchecked-type-assertion` — unchecked type assertions
- `revive/unhandled-error` — unhandled errors
- `revive/unreachable-code` — unreachable code after return
- `revive/use-errors-new` — use errors.New instead of fmt.Errorf
- `revive/use-fmt-print` — use fmt.Print instead of custom
- `rowserrcheck` — SQL rows.Err check
- `sqlclosecheck` — SQL rows/stmt close check
- `staticcheck/QF1005` — expand call to fmt.Sprintf
- `staticcheck/S1028` — simplify error construction
- `staticcheck/S1029` — simplify string conversion
- `staticcheck/SA1001` — discard result of fmt.Sprintf
- `staticcheck/SA1006` — printf with non-fatal format
- `staticcheck/SA1008` — invalid format in time.Parse
- `staticcheck/SA4013` — redundant return statement
- `staticcheck/SA4017` — unused variable or constant
- `staticcheck/SA4027` — redundant return after error
- `staticcheck/SA4032` — unreachable code
- `staticcheck/SA4033` — unreachable code after return
- `staticcheck/SA5001` — early return in defer
- `staticcheck/SA5009` — same pointer returned from function
- `staticcheck/SA6006` — missing fmt.Sprintf arguments
- `staticcheck/SA9008` — unreachable code in if-else chain
- `staticcheck/ST1004` — error naming convention
- `staticcheck/ST1008` — missing comment on exported function
- `wrapcheck` — wrapped error returns

### Missing Relationships

- `forcetypeassert` is in a different cluster but should be cross-referenced with `revive/unchecked-type-assertion` and `errorlint/asserts` — all deal with type assertion safety.
- `govet/errorsas` is in Dead Code Detection but is semantically related to `errorlint/asserts` — both deal with error type assertions.
- `revive/use-errors-new` is related to `staticcheck/S1028` (simplify error construction) but neither references the other.
- `nakedret` and `nonamedreturns` are in Style cluster but have meaningful error-handling overlap (naked returns often swallow errors).

### Surprising Connections

- `gosec/G104` (audit errors not checked) is the #1 god node with 66 edges, confirming it as the central hub connecting security and error handling concerns.
- `exhaustive` and `gochecksumtype` are grouped here because both deal with exhaustive pattern matching — a common source of unhandled cases that can cause runtime errors.

### Curation Guidance

- `guides/errcheck.md` `<related>` currently lists `err113, errname, wrapcheck, govet` — appropriate. Consider adding `nilerr` (returns nil after error check) and `rowserrcheck` (specific SQL case of unchecked errors).
- `guides/wrapcheck.md` `<related>` is empty. Add `errcheck` and `err113` — these form the core error handling trio.
- `guides/nilerr.md` `<related>` is empty. Add `errcheck` and `nilnesserr` — all three deal with nil-error antipatterns.
- `guides/nilnesserr.md` `<related>` is empty. Add `nilerr` and `errcheck`.
- `guides/nilnil.md` `<related>` is empty. Add `nilerr` — both deal with problematic nil returns.
- `guides/noinlineerr.md` `<related>` is empty. This linter is niche; consider leaving empty.
- `guides/rowserrcheck.md` `<related>` lists `sqlclosecheck, errcheck, bodyclose` — appropriate. Keep.
- `guides/sqlclosecheck.md` `<related>` lists `rowserrcheck, bodyclose, errcheck` — appropriate. Keep.
- `guides/errorlint/errorf.md` — Add `modernize/errorf` to `<related>` — same pattern, different era.
- `guides/errorlint/asserts.md` — Add `govet/errorsas` to `<related>` — both check error assertion correctness.
- `guides/errorlint/comparison.md` — Add `err113` to `<related>` — both deal with error comparison patterns.
- `guides/revive/unhandled-error.md` `<related>` is empty. Add `errcheck` — same problem domain.
- `guides/revive/error-strings.md` `<related>` is empty. Add `errname` — both deal with error naming conventions.
- `guides/staticcheck/SA4017.md` is a god node (37 edges). Its `<related>` lists `errcheck, SA4027` — appropriate.
- `guides/staticcheck/SA4027.md` `<related>` lists `errcheck, SA4017` — appropriate.

---

## Cluster: Concurrency Safety

**Cohesion:** 0.17
**Members:** 56 linters

### Member List

- `bodyclose` — unclosed HTTP response body
- `containedctx` — context.Context in struct
- `contextcheck` — missing context.Context propagation
- `gocritic/badLock` — inconsistent lock usage
- `gocritic/badSyncOnceFunc` — incorrect sync.Once usage
- `gocritic/deferInLoop` — defer in loop
- `gocritic/exitAfterDefer` — exit after defer
- `gocritic/syncMapLoadAndDelete` — sync.Map misuse
- `gocritic/unnecessaryDefer` — unnecessary defer statement
- `gosec/G118` — race condition via defer
- `gosec/G122` — missing mutex unlock
- `govet/atomic` — invalid atomic operations
- `govet/defers` — incorrect defer usage
- `govet/loopclosure` — captured loop variable in goroutine
- `govet/lostcancel` — lost context.CancelFunc
- `govet/sigchanyzer` — signal channel misuse
- `govet/testinggoroutine` — goroutine in test
- `govet/tests` — test function signature issues
- `govet/waitgroup` — sync.WaitGroup misuse
- `makezero` — slice append without initial allocation
- `revive/call-to-gc` — explicit GC call
- `revive/datarace` — potential data race
- `revive/defer` — defer in loop concerns
- `revive/forbidden-call-in-wg-go` — forbidden call in WaitGroup goroutine
- `revive/range-val-in-closure` — range value captured in closure
- `revive/time-equal` — time.Sub vs time.Equal
- `revive/use-any` — use any instead of interface{}
- `revive/use-waitgroup-go` — proper WaitGroup usage
- `revive/waitgroup-by-value` — WaitGroup passed by value
- `spancheck` — OpenTelemetry span closure check
- `staticcheck/QF1004` — unnecessary string conversion
- `staticcheck/S1005` — unnecessary blank assignment
- `staticcheck/SA1002` — invalid argument to regexp
- `staticcheck/SA1004` — suspicious channel usage
- `staticcheck/SA1005` — unnecessary blank import
- `staticcheck/SA1012` — nil context passed
- `staticcheck/SA1013` — context.Context not first parameter
- `staticcheck/SA1015` — using time.Tick in leaky way
- `staticcheck/SA1016` — time.Timer not stopped
- `staticcheck/SA1017` — channel already closed
- `staticcheck/SA1025` — non-unique interface method names
- `staticcheck/SA1029` — invalid key type in context.WithValue
- `staticcheck/SA1032` — invalid argument to time.Duration
- `staticcheck/SA2000` — sync.WaitGroup Add called in goroutine
- `staticcheck/SA2001` — empty critical section
- `staticcheck/SA2002` — called concurrently but not thread-safe
- `staticcheck/SA2003` — deferred close in loop
- `staticcheck/SA3000` — testmain not calling os.Exit
- `staticcheck/SA3001` — assigning to b.N in benchmark
- `staticcheck/SA4004` — suspicious LHS of assignment
- `staticcheck/SA4024` — identical expressions on both sides
- `staticcheck/SA5003` — unreachable code in defer
- `staticcheck/SA5004` — empty branch
- `staticcheck/SA6002` — storing non-pointer in sync.Pool
- `staticcheck/SA9001` — empty branch
- `staticcheck/SA9009` — same context passed twice

### Missing Relationships

- `govet/copylocks` is in Performance cluster but is semantically related to `govet/atomic`, `govet/waitgroup`, and `gocritic/badLock` — all deal with lock/mutex safety.
- `containedctx` is correctly placed but should cross-reference `contextcheck` (both deal with context.Context patterns).
- `makezero` is correctly placed but should cross-reference `govet/appends` (both deal with slice initialization).

### Surprising Connections

- `spancheck` (OpenTelemetry) is grouped with concurrency because spans must be closed properly — similar pattern to `bodyclose` (HTTP response) and `sqlclosecheck` (SQL rows). Resource cleanup is a cross-cutting concern.
- `staticcheck/SA1012` (nil context) and `staticcheck/SA1013` (context not first param) are in this cluster, not Error Handling, because they primarily cause goroutine cancellation failures.

### Curation Guidance

- `guides/bodyclose.md` `<related>` is empty. Add `contextcheck` and `noctx` — all deal with HTTP connection safety.
- `guides/contextcheck.md` `<related>` currently lists `revive, gocritic` — too vague. Replace with `noctx, bodyclose, govet/lostcancel` — specific and actionable.
- `guides/makezero.md` `<related>` is empty. Add `prealloc` — both deal with slice allocation optimization.
- `guides/spancheck.md` `<related>` currently lists `errcheck, contextcheck` — appropriate. Keep.
- `guides/containedctx.md` `<related>` currently lists `contextcheck, revive, gocritic` — too vague. Replace `revive` with `staticcheck/SA1013` (context not first param).
- `guides/govet/lostcancel.md` `<related>` lists `defers, httpresponse` — appropriate. Consider adding `contextcheck`.
- `guides/govet/loopclosure.md` `<related>` lists `testinggoroutine, defers` — appropriate. Add `revive/range-val-in-closure` — same pattern, different linter.
- `guides/govet/waitgroup.md` `<related>` lists `copylocks, atomic` — appropriate. Add `revive/waitgroup-by-value` — same concern.
- `guides/govet/atomic.md` `<related>` lists `copylocks, defers` — appropriate. Consider adding `gocritic/badLock`.

---

## Cluster: Style and Formatting

**Cohesion:** 0.21
**Members:** 50 linters

### Member List

- `asciicheck` — non-ASCII identifiers
- `decorder` — declaration order
- `dupword` — duplicate words in comments
- `forbidigo` — forbidden identifiers
- `funcorder` — function declaration order
- `gocheckcompilerdirectives` — compiler directive format
- `gochecknoglobals` — no global variables
- `gochecknoinits` — no init functions
- `goconst` — repeated strings should be constants
- `gocritic/codegenComment` — codegen comment format
- `gocritic/commentFormatting` — comment formatting
- `gocritic/commentedOutCode` — commented-out code
- `gocritic/commentedOutImport` — commented-out imports
- `gocritic/deprecatedComment` — deprecated comment format
- `gocritic/docStub` — stub documentation
- `gocritic/dupImport` — duplicate imports
- `gocritic/emptyDecl` — empty declarations
- `gocritic/todoCommentWithoutDetail` — TODO without detail
- `gocritic/whyNoLint` — nolint without explanation
- `godoclint` — documentation lint
- `godot` — comment period enforcement
- `godox` — TODO/FIXME/BUG detection
- `gofmt` — code formatting
- `gofumpt` — strict formatting
- `goheader` — file header enforcement
- `gomoddirectives` — go.mod directive format
- `gomodguard` — blocked module imports
- `gosmopolitan` — internationalization checks
- `importas` — import alias enforcement
- `lll` — line length limit
- `loggercheck` — logger usage patterns
- `misspell` — spelling mistakes
- `mnd` — magic numbers
- `nakedret` — naked returns
- `nlreturn` — blank line before return
- `nolintlint` — nolint directive format
- `nonamedreturns` — named return parameters
- `predeclared` — predeclared identifier shadowing
- `promlinter` — Prometheus metric naming
- `reassign` — reassignment of package vars
- `revive/comment-spacings` — comment spacing
- `revive/file-length-limit` — file length
- `revive/imports-blocklist` — blocked imports
- `revive/line-length-limit` — line length
- `sloglint` — slog usage patterns
- `usestdlibvars` — standard library variable usage
- `varnamelen` — variable name length
- `whitespace` — whitespace enforcement
- `wsl_v5` — whitespace style
- `zerologlint` — zerolog usage patterns

### Missing Relationships

- `goimports` is missing from this cluster — it should be here alongside `gofmt` and `gofumpt`.
- `revive/exported` and `revive/package-comments` are in other clusters but belong in this style cluster — they deal with documentation conventions.
- `gocritic/typeDefFirst` is in Code Complexity but could cross-reference `decorder` — both deal with declaration order.

### Surprising Connections

- `loggercheck`, `sloglint`, and `zerologlint` form a logging sub-cluster within Style — they share formatting and naming conventions for logging calls.
- `gosmopolitan` (internationalization) is grouped here because it shares formatting concern with `asciicheck` (both deal with character encoding in identifiers).

### Curation Guidance

- `guides/gofmt.md` `<related>` lists `gofumpt, govet, whitespace` — appropriate. Keep.
- `guides/gofumpt.md` `<related>` lists `whitespace, nlreturn, gofmt, govet` — appropriate. Keep.
- `guides/godot.md` `<related>` lists `dupword, godoclint, godox` — appropriate. Keep.
- `guides/decorder.md` `<related>` lists `gofmt, gofumpt, godoclint` — appropriate. Keep.
- `guides/wsl_v5.md` `<related>` lists `whitespace, nlreturn, nakedret` — appropriate. Keep.
- `guides/whitespace.md` `<related>` lists `wsl_v5, nlreturn, decorder` — appropriate. Keep.
- `guides/lll.md` `<related>` lists `funlen, godoclint, revive` — the `revive` reference is too vague. Replace with `revive/line-length-limit` for specificity.
- `guides/misspell.md` `<related>` lists `godot, lll, goheader` — appropriate (all deal with text quality).
- `guides/nakedret.md` `<related>` lists `nonamedreturns, nlreturn, errname` — appropriate. Consider adding `errcheck` (naked returns often hide errors).
- `guides/nonamedreturns.md` `<related>` lists `nakedret, nlreturn, errname` — appropriate.
- `guides/goconst.md` `<related>` lists `mnd, gochecknoglobals, dupl` — appropriate. Both `goconst` and `mnd` deal with magic values.
- `guides/sloglint.md` `<related>` lists `loggercheck, godot` — appropriate. Consider adding `zerologlint` — same logging domain.
- `guides/zerologlint.md` `<related>` lists `loggercheck, sloglint` — appropriate. Keep.
- `guides/gosmopolitan.md` `<related>` lists `predeclared, asciicheck, godot` — appropriate.
- `guides/loggercheck.md` `<related>` is empty. Add `sloglint` and `zerologlint` — logging linter trio.

---

## Cluster: Dead Code Detection

**Cohesion:** 0.21
**Members:** 40 linters

### Member List

- `gocritic/argOrder` — suspicious argument order
- `gocritic/caseOrder` — switch case ordering
- `gocritic/dupArg` — duplicate arguments
- `gocritic/dupBranchBody` — duplicate branch bodies
- `gocritic/dupCase` — duplicate case statements
- `gocritic/dupOption` — duplicate option values
- `gocritic/dupSubExpr` — duplicate sub-expressions
- `gocritic/evalOrder` — evaluation order dependency
- `gocritic/mapKey` — duplicate map key literals
- `gosec/G709` — integer overflow
- `govet/assign` — useless assignment
- `govet/bools` — redundant boolean expressions
- `govet/composites` — unkeyed composite literals
- `govet/errorsas` — errors.As target type
- `govet/nilfunc` — nil function comparison
- `govet/shift` — invalid shift amount
- `govet/slog` — slog argument types
- `govet/stdmethods` — standard method signatures
- `govet/stringintconv` — string/int conversion
- `govet/structtag` — struct tag format
- `govet/timeformat` — time format string
- `govet/unmarshal` — unmarshal argument type
- `govet/unreachable` — unreachable code
- `staticcheck/SA4000` — duplicate blank identifiers
- `staticcheck/SA4005` — unnecessary field assignment
- `staticcheck/SA4006` — value assigned but unused
- `staticcheck/SA4008` — variable assigned but unused
- `staticcheck/SA4009` — argument overwritten before use
- `staticcheck/SA4010` — useless break in switch
- `staticcheck/SA4016` — redundant type assertion
- `staticcheck/SA4018` — self-assignment
- `staticcheck/SA4019` — duplicate if-else condition
- `staticcheck/SA5000` — nil pointer dereference in assignment
- `staticcheck/SA5002` — empty slice append
- `staticcheck/SA5011` — nil pointer dereference
- `staticcheck/SA6001` — MapIter.Next called without Key/Value
- `staticcheck/SA9006` — dubious bit shifting
- `staticcheck/SA9007` — XML element or attribute name
- `unparam` — unused function parameters
- `unused` — unused declarations

### Missing Relationships

- `govet/errorsas` is in this cluster but is semantically related to `errorlint/asserts` in Error Handling. Cross-reference needed.
- `staticcheck/SA4006` and `staticcheck/SA4008` (both "assigned but unused") overlap with `ineffassign` in Performance. Consider cross-referencing.
- `wastedassign` is in Performance but is semantically identical to `staticcheck/SA4006` — they should reference each other.

### Surprising Connections

- `govet/errorsas` bridges this cluster to Error Handling — it checks whether errors.As targets are valid, which is both a code quality issue and an error-handling concern.

### Curation Guidance

- `guides/unused.md` `<related>` is empty. Add `unparam` — both detect dead code. This is a high-priority curation.
- `guides/unparam.md` `<related>` is empty. Add `unused` — reciprocal reference.
- `guides/govet/unreachable.md` `<related>` lists `tests, defers` — appropriate. Consider adding `staticcheck/SA4032` — both detect unreachable code.
- `guides/govet/errorsas.md` `<related>` lists `ifaceassert, nilfunc` — appropriate. Add `errorlint/asserts` — cross-cluster bridge.
- `guides/staticcheck/SA4006.md` `<related>` lists `SA4016, SA4015` — appropriate. Consider adding `ineffassign` — same domain (unused assignments).
- `guides/staticcheck/SA4019.md` `<related>` lists `SA4000, SA4018` — appropriate. Keep.
- `guides/gocritic/dupArg.md` `<related>` lists `dupSubExpr, dupBranchBody, argOrder` — excellent. Keep.
- `guides/gocritic/dupBranchBody.md` `<related>` lists `dupCase, dupArg, dupSubExpr` — excellent. Keep.

---

## Cluster: Code Complexity

**Cohesion:** 0.21
**Members:** 33 linters

### Member List

- `arangolint` — ArangoDB query linting
- `cyclop` — cyclomatic complexity
- `dogsled` — excessive blank assignments
- `dupl` — code duplication detection
- `errchkjson` — JSON encoding error check
- `funlen` — function length
- `gocognit` — cognitive complexity
- `gocritic/defaultCaseOrder` — default case positioning
- `gocritic/elseif` — else-if chain
- `gocritic/emptyFallthrough` — empty fallthrough
- `gocritic/ifElseChain` — if-else chain length
- `gocritic/initClause` — init clause in if
- `gocritic/nestingReduce` — nesting reduction
- `gocritic/paramTypeCombine` — parameter type combining
- `gocritic/singleCaseSwitch` — single-case switch
- `gocritic/switchTrue` — switch true
- `gocritic/tooManyResultsChecker` — too many return values
- `gocritic/typeAssertChain` — type assertion chain
- `gocritic/typeDefFirst` — type definition before usage
- `gocritic/typeSwitchVar` — type switch variable
- `gocritic/unlabelStmt` — unnecessary label
- `gocritic/unnamedResult` — unnamed return values
- `gocritic/unnecessaryBlock` — unnecessary block
- `gocyclo` — cyclomatic complexity (alternate)
- `interfacebloat` — interface method count
- `ireturn` — interface return types
- `maintidx` — maintainability index
- `nestif` — nesting level
- `revive/cognitive-complexity` — cognitive complexity (revive)
- `revive/cyclomatic` — cyclomatic complexity (revive)
- `revive/function-length` — function length (revive)
- `revive/max-control-nesting` — max nesting (revive)
- `staticcheck/ST1016` — consistent pointer type

### Missing Relationships

- `gocognit` and `gocyclo` are correctly cross-referenced with `cyclop` and `maintidx`. The complexity linter group is well-connected.
- `revive/cognitive-complexity`, `revive/cyclomatic`, `revive/function-length`, and `revive/max-control-nesting` should all cross-reference each other AND the standalone linters (`cyclop`, `gocognit`, `gocyclo`, `maintidx`, `nestif`).
- `interfacebloat` overlaps with `funlen` — both measure "too much stuff" in a construct.

### Surprising Connections

- `arangolint` (ArangoDB query linting) appears here because graphify detected semantic overlap with complexity concerns — ArangoDB queries with complex filter chains mirror Go code complexity patterns.
- `errchkjson` appears here because JSON encoding error handling is often a source of complexity (nested error checks).

### Curation Guidance

- `guides/cyclop.md` `<related>` lists `gocyclo, gocognit, maintidx, funlen, nestif` — comprehensive. Keep as-is.
- `guides/gocyclo.md` `<related>` lists `cyclop, gocognit, maintidx, funlen` — appropriate. Consider adding `revive/cyclomatic` — same metric, different linter.
- `guides/funlen.md` `<related>` lists `cyclop, gocyclo, gocognit, nestif` — appropriate. Consider adding `revive/function-length` — same concern, different linter.
- `guides/nestif.md` `<related>` lists `gocognit, cyclop, gocyclo, funlen` — appropriate. Consider adding `revive/max-control-nesting` — same nesting concern.
- `guides/maintidx.md` `<related>` lists `gocyclo, gocognit, cyclop, funlen` — appropriate. Keep.
- `guides/gocognit.md` `<related>` lists `gocyclo, cyclop, maintidx, nestif` — appropriate. Consider adding `revive/cognitive-complexity` — same metric.
- `guides/interfacebloat.md` `<related>` lists `funlen, gocyclo, revive` — too vague. Replace `revive` with `ireturn` — both deal with interface design complexity.
- `guides/ireturn.md` `<related>` lists `interfacebloat, revive, gocritic` — too vague. Replace `revive` with `revive/max-public-structs` and `gocritic` with `gocritic/unnamedResult`.
- `guides/dupl.md` `<related>` lists `funlen, gocyclo, gocognit` — appropriate (duplicated code increases complexity).
- `guides/gocritic/nestingReduce.md` `<related>` lists `elseif, ifElseChain, unnecessaryBlock` — excellent. Keep.
- `guides/gocritic/ifElseChain.md` `<related>` lists `elseif, singleCaseSwitch, switchTrue` — excellent. Keep.
- `guides/dogsled.md` `<related>` lists `unparam, errcheck, funlen` — the `unparam` and `errcheck` connections are weak. Consider replacing with `gocritic/tooManyResultsChecker` — both deal with excessive return values.

---

## Cluster: Performance Optimization

**Cohesion:** 0.25
**Members:** 29 linters

### Member List

- `fatcontext` — context struct size
- `gocritic/appendAssign` — append result assignment
- `gocritic/appendCombine` — combined append calls
- `gocritic/badCond` — suspicious conditions
- `gocritic/badSorting` — incorrect sort usage
- `gocritic/builtinShadow` — builtin shadowing
- `gocritic/builtinShadowDecl` — builtin shadowing in declarations
- `gocritic/hugeParam` — large parameter passed by value
- `gocritic/indexAlloc` — unnecessary allocation in index
- `gocritic/offBy1` — off-by-one errors
- `gocritic/rangeAppendAll` — append entire range
- `gocritic/rangeExprCopy` — range expression copy
- `gocritic/rangeValCopy` — range value copy
- `gocritic/sloppyLen` — sloppy len usage
- `gocritic/sloppyReassign` — sloppy reassignment
- `gocritic/sortSlice` — sort.Slice misuse
- `gocritic/truncateCmp` — truncation in comparison
- `govet/appends` — append result unused
- `govet/copylocks` — copying lock values
- `govet/hostport` — host:port splitting
- `ineffassign` — ineffective assignments
- `nosprintfhostport` — sprintf for host:port
- `perfsprint` — performance-aware fmt.Sprintf
- `prealloc` — slice preallocation
- `revive/atomic` — atomic usage patterns
- `staticcheck/SA1027` — atomic argument type
- `staticcheck/SA1030` — strconv pattern simplification
- `unconvert` — unnecessary type conversions
- `wastedassign` — wasted assignments

### Missing Relationships

- `staticcheck/SA4006` (assigned but unused) is in Dead Code Detection but is semantically identical to `wastedassign` — they should cross-reference.
- `govet/copylocks` should cross-reference `govet/atomic` and `gocritic/badLock` in Concurrency — copying locks is a concurrency safety issue too.

### Surprising Connections

- `fatcontext` (context struct size) is grouped with performance because large contexts consume memory — but it also relates to `containedctx` (context in struct) in Concurrency. A cross-cluster reference is warranted.
- `perfsprint` and `nosprintfhostport` form a sprint-related sub-cluster — both optimize fmt.Sprintf usage.

### Curation Guidance

- `guides/prealloc.md` `<related>` is empty. Add `wastedassign` and `ineffassign` — all detect wasted allocation/assignment patterns. This is high-priority.
- `guides/ineffassign.md` `<related>` is empty. Add `wastedassign` and `prealloc` — same domain.
- `guides/wastedassign.md` `<related>` is empty. Add `ineffassign` and `prealloc` — reciprocal.
- `guides/unconvert.md` `<related>` is empty. Add `perfsprint` — both eliminate unnecessary operations.
- `guides/perfsprint.md` `<related>` lists `nosprintfhostport, govet, errcheck` — replace `govet` with `govet/printf` for specificity.
- `guides/nosprintfhostport.md` `<related>` lists `perfsprint, govet, errcheck` — replace `govet` with `govet/hostport` for specificity.
- `guides/govet/copylocks.md` `<related>` lists `atomic, waitgroup` — appropriate. Add `gocritic/badLock` — same lock concern.
- `guides/fatcontext.md` `<related>` is empty. Add `gocritic/hugeParam` and `gocritic/rangeValCopy` — all deal with unnecessary data copying.
- `guides/gocritic/hugeParam.md` `<related>` is empty. Add `gocritic/rangeValCopy` and `prealloc` — same optimization domain.
- `guides/gocritic/rangeValCopy.md` `<related>` is empty. Add `gocritic/rangeExprCopy` and `gocritic/hugeParam` — same copy concern.

---

## Cluster: Code Simplification

**Cohesion:** 0.09
**Members:** 27 linters

### Member List

- `gocritic/assignOp` — assignment operation simplification
- `gocritic/boolExprSimplify` — boolean expression simplification
- `gocritic/captLocal` — capitalize local naming
- `gocritic/deferUnlambda` — defer without lambda
- `gocritic/emptyStringTest` — empty string test pattern
- `gocritic/exposedSyncMutex` — exposed sync.Mutex field
- `gocritic/filepathJoin` — filepath.Join usage
- `gocritic/hexLiteral` — hex literal format
- `gocritic/httpNoBody` — http.NoBody usage
- `gocritic/methodExprCall` — method expression call
- `gocritic/newDeref` — new(T) dereference
- `gocritic/octalLiteral` — octal literal format
- `gocritic/preferFilepathJoin` — prefer filepath.Join
- `gocritic/ptrToRefParam` — pointer to reference parameter
- `gocritic/redundantSprint` — redundant fmt.Sprint
- `gocritic/stringConcatSimplify` — string concatenation simplification
- `gocritic/stringsCompare` — strings.Compare simplification
- `gocritic/timeExprSimplify` — time expression simplification
- `gocritic/typeUnparen` — unnecessary type parentheses
- `gocritic/underef` — unnecessary dereference
- `gocritic/unlambda` — lambda simplification
- `gocritic/unslice` — unnecessary slicing
- `gocritic/valSwap` — value swap pattern
- `gocritic/wrapperFunc` — wrapper function simplification
- `gocritic/yodaStyleExpr` — yoda-style expression
- `staticcheck/SA4021` — unnecessary type assertion
- `staticcheck/ST1022` — documentation format

### Missing Relationships

- This cluster is almost entirely gocritic rules — they naturally group together as "code simplification patterns". No significant missing relationships.
- `gocritic/preferFilepathJoin` and `gocritic/filepathJoin` should reference each other (they currently do).

### Surprising Connections

- Low cohesion (0.09) indicates these are loosely connected — they share the "simplification" theme but individually address different code patterns. This is expected for gocritic's broad rule set.

### Curation Guidance

- `guides/gocritic/assignOp.md` `<related>` lists `stringConcatSimplify, yodaStyleExpr` — appropriate. Keep.
- `guides/gocritic/boolExprSimplify.md` `<related>` lists `yodaStyleExpr, elseif` — appropriate. Keep.
- `guides/gocritic/deferUnlambda.md` `<related>` lists `unlambda, unnecessaryBlock` — appropriate. Keep.
- `guides/gocritic/redundantSprint.md` `<related>` lists `stringConcatSimplify, wrapperFunc` — appropriate. Keep.
- `guides/gocritic/underef.md` `<related>` lists `newDeref, typeUnparen` — appropriate. Keep.
- `guides/gocritic/unlambda.md` `<related>` is empty. Add `gocritic/deferUnlambda` — same lambda simplification pattern.
- `guides/gocritic/unslice.md` `<related>` is empty. Add `gocritic/typeUnparen` — both remove unnecessary syntax.
- All other gocritic rules in this cluster have appropriate existing `<related>` tags. No changes needed.

---

## Cluster: Deprecated API Patterns

**Cohesion:** 0.10
**Members:** 24 linters

### Member List

- `depguard` — dependency guard
- `staticcheck/SA1000` — invalid regex pattern
- `staticcheck/SA1010` — time.Sleep duration
- `staticcheck/SA1019` — deprecated API usage
- `staticcheck/SA4001` — duplicate test in if/else
- `staticcheck/SA4003` — unsigned comparison
- `staticcheck/SA4011` — useless break in switch
- `staticcheck/SA4012` — empty branch
- `staticcheck/SA4014` — duplicate case
- `staticcheck/SA4015` — unnecessary type assertion
- `staticcheck/SA4022` — nil pointer dereference in select
- `staticcheck/SA4025` — unnecessary type conversion
- `staticcheck/SA4026` — unnecessary break
- `staticcheck/SA4028` — empty default case
- `staticcheck/SA4030` — infinite recursive call
- `staticcheck/SA5005` — empty struct field
- `staticcheck/SA5012` — passing odd-sized struct
- `staticcheck/SA6000` — MapIter misuse
- `staticcheck/SA6003` — string concatenation in loop
- `staticcheck/SA6004` — error variable shadowed
- `staticcheck/SA6005` — inefficient string comparison
- `staticcheck/SA9002` — invalid file mode
- `staticcheck/SA9003` — empty range
- `staticcheck/SA9004` — omitted JSON tags

### Missing Relationships

- This cluster groups staticcheck rules that detect patterns that are either deprecated or commonly replaced by modern Go patterns. The cluster name is somewhat misleading — it's really "Static Analysis Miscellaneous" rather than strictly "Deprecated APIs".
- `staticcheck/SA1019` (deprecated API usage) should cross-reference `gocritic/deprecatedComment` — both deal with deprecation.

### Surprising Connections

- `depguard` appears in this cluster because it guards against deprecated dependencies — a natural fit with staticcheck's deprecated API detection.

### Curation Guidance

- `guides/staticcheck/SA1019.md` `<related>` lists `depguard` — appropriate (both block deprecated dependencies). Keep.
- `guides/staticcheck/SA1000.md` `<related>` lists `SA1010, gosec/G204` — the gosec/G204 reference seems noise. Consider removing.
- `guides/staticcheck/SA4014.md` `<related>` lists `SA4000, SA4001, SA4003` — appropriate (duplicate/unused pattern group). Keep.
- `guides/staticcheck/SA4012.md` `<related>` lists `SA4014, SA4026` — appropriate. Keep.
- `guides/staticcheck/SA6005.md` `<related>` lists `SA6000, SA6006` — appropriate. Keep.
- Most staticcheck rules in this cluster already cross-reference each other. The existing `<related>` network is dense and appropriate.

---

## Cluster: Testing Frameworks

**Cohesion:** 0.92
**Members:** 18 linters

### Member List

- `ginkgolinter/compare-assertion` — Ginkgo comparison assertion patterns
- `ginkgolinter/error-assertion` — Ginkgo error assertion patterns
- `ginkgolinter/nil-assertion` — Ginkgo nil assertion patterns
- `ginkgolinter/succeed-matcher` — Ginkgo succeed matcher patterns
- `paralleltest` — missing t.Parallel in tests
- `testableexamples` — testable examples
- `testifylint/bool-compare` — testify bool comparison patterns
- `testifylint/error-as` — testify error assertion
- `testifylint/error-nil` — testify nil error assertion
- `testifylint/expected-actual` — testify expected/actual argument order
- `testifylint/go-require` — testify require in goroutines
- `testifylint/nil-compare` — testify nil comparison patterns
- `testifylint/require-error` — testify require.Error patterns
- `testifylint/useless-assert` — testify useless assertions
- `testpackage` — test package naming
- `thelper` — test helper function patterns
- `tparallel` — test parallel patterns
- `usetesting` — use testing package patterns

### Missing Relationships

- `ginkgolinter/async-assertion`, `ginkgolinter/focus-container`, `ginkgolinter/have-len-zero`, `ginkgolinter/len-assertion`, `ginkgolinter/spec-pollution`, `ginkgolinter/type-compare`, and `ginkgolinter/expect-to` are in other clusters. They should be in this Testing cluster. This is a graphify clustering artifact — these ginkgolinter rules were not strongly connected enough to the main group.
- `testifylint/float-compare`, `testifylint/suite-method-signature`, `testifylint/suite-dont-use-pkg`, `testifylint/suite-extra-assert-call`, `testifylint/suite-thelper`, `testifylint/contains-unnecessary-format`, `testifylint/suite-broken-parallel`, `testifylint/formatter`, `testifylint/blank-import`, and `testifylint/empty` are also missing from this cluster.
- `govet/tests` and `govet/testinggoroutine` are in Concurrency but should cross-reference this cluster.

### Surprising Connections

- Highest cohesion score (0.92) — testing linters naturally form the tightest cluster. This confirms that testifylint and ginkgolinter rules have strong internal relationships.
- `usetesting` is here because it promotes using `testing` package helpers — a meta-concern that applies to all test frameworks.

### Curation Guidance

- `guides/paralleltest.md` `<related>` lists `tparallel, thelper, testpackage` — appropriate. Consider adding `usetesting`.
- `guides/tparallel.md` `<related>` lists `paralleltest, thelper, testpackage` — appropriate. Consider adding `usetesting`.
- `guides/thelper.md` `<related>` lists `paralleltest, tparallel, testpackage` — appropriate. Consider adding `usetesting` and `testifylint/suite-thelper`.
- `guides/testpackage.md` `<related>` lists `paralleltest, thelper, testableexamples` — appropriate. Keep.
- `guides/testableexamples.md` `<related>` lists `testpackage, godoclint, thelper` — appropriate. Keep.
- `guides/usetesting.md` `<related>` lists `thelper, testpackage, paralleltest` — appropriate. Keep.
- `guides/testifylint/error-as.md` `<related>` is empty. Add `testifylint/error-nil` and `errorlint/asserts` — both deal with error assertion patterns.
- `guides/testifylint/error-nil.md` `<related>` is empty. Add `testifylint/error-as` and `errcheck` — nil error checks relate to general error checking.
- `guides/testifylint/require-error.md` `<related>` is empty. Add `testifylint/error-nil` — both check error assertions.
- `guides/testifylint/expected-actual.md` `<related>` is empty. Add `testifylint/bool-compare` and `testifylint/nil-compare` — all deal with assertion argument order.
- `guides/testifylint/useless-assert.md` `<related>` is empty. Add `testifylint/expected-actual` — both detect redundant assertions.
- `guides/testifylint/nil-compare.md` `<related>` is empty. Add `testifylint/bool-compare` — both deal with comparison patterns.
- `guides/ginkgolinter/error-assertion.md` `<related>` is empty. Add `ginkgolinter/nil-assertion` — both are common Ginkgo assertion patterns.
- `guides/ginkgolinter/nil-assertion.md` `<related>` is empty. Add `ginkgolinter/error-assertion` — reciprocal.
- `guides/ginkgolinter/compare-assertion.md` `<related>` is empty. Add `ginkgolinter/nil-assertion` — related assertion patterns.
- `guides/ginkgolinter/succeed-matcher.md` `<related>` is empty. Add `ginkgolinter/error-assertion` — both deal with error checking in Ginkgo.

---

## Cross-Cluster Analysis

### Overlapping Concerns

1. **Error Handling ↔ Security**: `gosec/G104` (errors not checked) is the #1 god node. Error handling and security overlap significantly — unchecked errors can lead to security vulnerabilities. The graph correctly identifies this bridge.

2. **Concurrency ↔ Error Handling**: `govet/lostcancel` (lost context.CancelFunc) sits in Concurrency but connects deeply to error handling — lost cancellations cause goroutine leaks and timeout failures. Cross-reference `contextcheck` ↔ `errcheck` is warranted.

3. **Dead Code ↔ Performance**: `wastedassign`, `ineffassign`, `unused`, and `unparam` form a dead-code/performance nexus. Unused code wastes both CPU cycles and developer attention. The graph splits them across two clusters — Phase 57 should ensure cross-references between `unused` ↔ `unparam` ↔ `wastedassign` ↔ `ineffassign`.

4. **Complexity ↔ Style**: Code complexity tools (`cyclop`, `funlen`) and style tools (`lll`, `revive/file-length-limit`) measure related but different aspects. Long functions are both complex and hard to format. The revive variants (`revive/cyclomatic`, `revive/function-length`) bridge these clusters.

### Prioritization for Phase 57

**High-priority curation (most impact, most guides affected):**

1. **Fill empty `<related>` tags in Error Handling**: `wrapcheck`, `nilerr`, `nilnesserr`, `nilnil`, `noinlineerr`, `revive/unhandled-error`, `revive/error-strings` — all have empty related tags despite being well-connected in the graph.

2. **Fill empty `<related>` tags in Performance**: `prealloc`, `ineffassign`, `wastedassign`, `unconvert`, `gocritic/hugeParam`, `gocritic/rangeValCopy` — all have empty related tags.

3. **Fill empty `<related>` tags in Testing Frameworks**: All testifylint and ginkgolinter rules except those in the main cluster have empty related tags. Add cross-references within each framework family.

4. **Fix vague `<related>` references**: Replace bare `revive`, `gocritic`, `govet` references with specific rule names. Examples:
   - `contextcheck` → replace `revive` with `noctx`
   - `lll` → replace `revive` with `revive/line-length-limit`
   - `perfsprint` → replace `govet` with `govet/printf`

5. **Connect the god nodes**: `errcheck` (37 edges) and `gosec/G104` (66 edges) are natural hubs. Any guide that deals with error returns or security should reference these.

**Medium-priority curation:**

6. **Cross-reference revive variants**: `revive/cyclomatic` ↔ `cyclop`, `revive/function-length` ↔ `funlen`, `revive/max-control-nesting` ↔ `nestif` — these are the same metrics implemented by different linters.

7. **Connect Concurrency cluster**: Add missing cross-references between `bodyclose`/`contextcheck`/`makezero` and their related govet/revive counterparts.

8. **Connect Code Simplification cluster**: Most gocritic rules already cross-reference, but some are empty (e.g., `gocritic/unlambda`, `gocritic/unslice`).

**Low-priority (existing tags are appropriate):**

9. **Security Auditing cluster**: The 61 gosec rules have dense internal cross-references. No major changes needed.
10. **Dead Code Detection cluster**: Most rules already cross-reference appropriately.
11. **Deprecated API Patterns cluster**: staticcheck rules already form a dense cross-reference network.

### Singleton and Small-Cluster Guides

187 guides form singletons (communities of 1) — these have no significant graph connections. For Phase 57:
- If a singleton has an empty `<related>` tag AND no meaningful semantic connections: leave the `<related>` tag empty (per project decision: omit/empty for trivial cases).
- If a singleton has an empty `<related>` tag but DOES have semantic connections to named clusters: add 1-2 relevant references.
- Notable singletons that need attention: `gosec/G118` (race condition via defer) and `gosec/G122` (missing mutex unlock) — both should cross-reference Concurrency cluster members.

---

*Generated by graphify pipeline for Phase 57 related-tag curation.*
*Data source: graphify-out/graph.json (629 nodes, 2,232 edges, 204 communities)*
