package environment

import (
	"os"
	"runtime"
	"strings"
	"time"
)

var ()

func IsCompiled() bool {
	if strings.HasPrefix(os.Args[0], "/var/folders/") ||
		strings.HasPrefix(os.Args[0], "/tmp/go-build") ||
		strings.Contains(os.Args[0], "\\AppData\\Local\\Temp\\") {
		return false
	}
	return true
}

func Compiler() string {
	return runtime.Compiler
}

func GOARCH() string {
	return runtime.GOARCH
}

func GOOS() string {
	return runtime.GOOS
}

func GOROOT() string {
	return runtime.GOROOT()
}

func GOVER() string {
	return runtime.Version()
}

func NumCPU() int {
	return runtime.NumCPU()
}

func GOPATH() string {
	return os.Getenv("GOPATH")
}

func GetFormattedTime() string {
	return Now("Monday, 2 Jan 2006")
}

func Now(layout string) string {
	return time.Now().Format(layout)
}
