# govet: framepointer

<instructions>
Reports assembly functions that clobber the frame pointer register. Go's runtime expects the frame pointer to be preserved for stack traces and profiling. Assembly functions must save and restore the frame pointer following the Go ABI conventions.

Follow Go's assembly conventions: save the frame pointer on entry and restore it on return.
</instructions>

<examples>
## Bad
```asm
TEXT ·broken(SB), NOSPLIT, $0
    MOVQ BP, 123(SP) // clobbers BP without saving
```

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
- Overwriting BP register without saving the old frame pointer
- Assembly function missing frame pointer setup
- Incorrect frame size for frame pointer storage
</patterns>

<related>
asmdecl, cgocall
</related>
