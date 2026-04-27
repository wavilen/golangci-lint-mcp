# govet: hostport

<instructions>
Detects misuse of `net.JoinHostPort` where the host argument already contains a port (e.g., passing `"example.com:8080"` as the host). `net.JoinHostPort` adds its own separator, producing a malformed address like `"example.com:8080:80"`.

Pass the host and port as separate arguments. If you already have a combined address, use it directly instead of calling `JoinHostPort`.
</instructions>

<examples>
## Good
```go
addr := net.JoinHostPort("example.com", "443")
// produces "example.com:443"
```
</examples>

<patterns>
- Split `host:port` strings before passing the host to `net.JoinHostPort`
- Avoid calling `net.JoinHostPort` on already-combined network addresses
- Separate host and port from hardcoded address strings before calling `JoinHostPort`
</patterns>

<related>
govet/httpresponse, nosprintfhostport
</related>
