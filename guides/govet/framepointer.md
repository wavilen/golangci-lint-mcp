# govet: framepointer

<instructions>
Reports assembly functions that clobber the frame pointer register. Go's runtime expects the frame pointer to be preserved for stack traces and profiling. Assembly functions must save and restore the frame pointer following the Go ABI conventions.

Follow Go's assembly conventions: save the frame pointer on entry and restore it on return.
</instructions>

<examples>
## Good
```asm
TEXT ·correct(SB), NOSPLIT, $16
    MOVQ BP, -8(SP)
    LEAQ -8(SP), BP
    // ... function body ...
    MOVQ -8(SP), BP
    RET
```
</examples>

<patterns>
- Save the old frame pointer before modifying the BP register in assembly functions
- Add frame pointer setup (`MOVQ`/`LEAQ` BP) in all assembly functions
- Allocate the correct frame size to accommodate frame pointer storage
</patterns>

<related>
govet/asmdecl, govet/cgocall
</related>
