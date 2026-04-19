# govet: sigchanyzer

<instructions>
Reports unbuffered channels used with `signal.Notify`. `signal.Notify` sends on the channel for each received signal, and an unbuffered channel can miss signals if the receiver is not ready. Use a buffered channel to avoid losing signals.
</instructions>

<examples>
## Bad
```go
ch := make(chan os.Signal)
signal.Notify(ch, syscall.SIGTERM)
// unbuffered — signal may be dropped if not immediately consumed
```

## Good
```go
ch := make(chan os.Signal, 1)
signal.Notify(ch, syscall.SIGTERM)
// buffered — signal is retained until consumed
```
</examples>

<patterns>
- `make(chan os.Signal)` passed to `signal.Notify`
- Unbuffered signal channel that can lose notifications
- Signal handler using blocking channel send
</patterns>

<related>
lostcancel, defers
</related>
