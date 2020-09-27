package system

import (
	"runtime"
	"strings"
)

var singletonIsLinux *bool

// CrossPlatformPath this function exist to solve problem
// when windows use "\\" and linux use "/"
func CrossPlatformPath(path string) string {
	if singletonIsLinux == nil {
		tempPlatform := isLinux()
		singletonIsLinux = &tempPlatform
	}
	if *singletonIsLinux {
		path = strings.ReplaceAll(path, "\\", "/")
	} else {
		path = strings.ReplaceAll(path, "/", "\\")
	}
	return path
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}
