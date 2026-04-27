# govet: asmdecl

<instructions>
Reports mismatches between assembly (`.s`) function declarations and their Go prototypes. Common issues include wrong argument sizes, missing return values, or incorrect function signatures in assembly files.

Fix the assembly function signatures to match the Go declarations exactly — argument offsets, sizes, and return value layout must align with the Go ABI.
</instructions>

<examples>
## Good
```go
// Assembly (correct: return value at ret+16(FP))
TEXT ·Add(SB), NOSPLIT, $0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    ADDQ BX, AX
    MOVQ AX, ret+16(FP)
    RET
```
</examples>

<patterns>
- Add return value storage in assembly functions — map return values to correct frame offsets
- Match argument sizes between Go declarations and assembly implementations
- Use correct frame pointer offsets in assembly to match the Go ABI
</patterns>

<related>
govet/framepointer, govet/cgocall
</related>
