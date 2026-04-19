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
- `fmt.Sprintf` with `%s:%d` or `%s:%s` for address construction
- String concatenation for host:port: `host + ":" + port`
- IPv6 addresses that break when formatted as `[host]:port` manually
- URL construction via string formatting instead of `url.URL`
</patterns>

<related>
perfsprint, govet
</related>
