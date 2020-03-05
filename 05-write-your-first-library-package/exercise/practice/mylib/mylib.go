package mylib

import (
	"fmt"
	"runtime"
)

// DisplayGoVersion ...
// Version returns the Go tree's version string.
// It is either the commit hash and date at the time of the build or,
// when possible, a release tag like "go1.3".
func DisplayGoVersion() {
	fmt.Println(runtime.Version())
}
