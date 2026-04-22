package version

import (
	"runtime/debug"
	"strings"
)

func init() {
	if Server != "dev" {
		return
	}
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, s := range bi.Settings {
		if s.Key == "vcs.revision" {
			Server = s.Value[:9]
			if len(s.Value) < 9 {
				Server = s.Value
			}
			break
		}
	}
	for _, s := range bi.Settings {
		if s.Key == "vcs.tag" && strings.HasPrefix(s.Value, "v") {
			Server = strings.TrimPrefix(s.Value, "v")
			break
		}
	}
}

var Server = "dev"
