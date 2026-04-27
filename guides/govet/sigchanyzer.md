# govet: sigchanyzer

<instructions>
Reports unbuffered channels used with `signal.Notify`. `signal.Notify` sends on the channel for each received signal, and an unbuffered channel can miss signals if the receiver is not ready. Use a buffered channel to avoid losing signals.
</instructions>

<examples>
## Good
```go
ch := make(chan os.Signal, 1)
signal.Notify(ch, syscall.SIGTERM)
// buffered — signal is retained until consumed
```
</examples>

<patterns>
- Use a buffered channel (`make(chan os.Signal, 1)`) with `signal.Notify`
- Create buffered signal channels to avoid losing notifications
- Use non-blocking or buffered sends in signal handlers
</patterns>

<related>
govet/lostcancel, govet/defers
</related>
