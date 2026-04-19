# govet: hostport

<instructions>
Detects misuse of `net.JoinHostPort` where the host argument already contains a port (e.g., passing `"example.com:8080"` as the host). `net.JoinHostPort` adds its own separator, producing a malformed address like `"example.com:8080:80"`.

Pass the host and port as separate arguments. If you already have a combined address, use it directly instead of calling `JoinHostPort`.
</instructions>

<examples>
## Bad
```go
addr := net.JoinHostPort("example.com:8080", "443")
// produces "example.com:8080:443" — malformed
```

## Good
```go
addr := net.JoinHostPort("example.com", "443")
// produces "example.com:443"
```
</examples>

<patterns>
- Passing `host:port` string as host argument to `net.JoinHostPort`
- Joining an already-combined network address
- Hardcoded host strings that include port numbers
</patterns>

<related>
httpresponse, nosprintfhostport
</related>
