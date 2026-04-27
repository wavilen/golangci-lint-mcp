package version

import (
	"runtime/debug"
	"strings"
)

const shortHashLen = 9

func init() {
	if Server != "dev" {
		return
	}
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, s := range buildInfo.Settings {
		if s.Key == "vcs.revision" {
			if len(s.Value) < shortHashLen {
				Server = s.Value
			} else {
				Server = s.Value[:shortHashLen]
			}
			break
		}
	}
	for _, s := range buildInfo.Settings {
		if s.Key == "vcs.tag" && strings.HasPrefix(s.Value, "v") {
			Server = strings.TrimPrefix(s.Value, "v")
			break
		}
	}
}

var Server = "dev"
