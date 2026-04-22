# nosprintfhostport

<instructions>
Nosprintfhostport detects uses of `fmt.Sprintf` to construct host:port strings. Using `fmt.Sprintf("%s:%d", host, port)` produces incorrect results when the host contains a port or is an IPv6 address.

Use `net.JoinHostPort` which correctly handles IPv6 addresses and other edge cases.
</instructions>

<examples>
## Bad
```go
addr := fmt.Sprintf("%s:%d", host, port)
```

## Good
```go
addr := net.JoinHostPort(host, strconv.Itoa(port))
```
</examples>

<patterns>
- Replace `fmt.Sprintf` address construction with `net.JoinHostPort`
- Use `net.JoinHostPort` instead of string concatenation for host:port
- Use `net.JoinHostPort` to correctly handle IPv6 address formatting
- Build URLs using `url.URL` struct instead of string formatting
</patterns>

<related>
perfsprint, govet
</related>
