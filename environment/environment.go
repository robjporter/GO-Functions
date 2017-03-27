package environment

import (
	"fmt"
	"github.com/robjporter/go-functions/as"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	Prefix = ""
)

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

func GetEnv(n string, def interface{}) interface{} {
	value := os.Getenv(prefixedName(n))
	if value == "" {
		return def
	}
	return value
}

//GetString returns a environment variable as string
func GetEnvString(n string, def string) string {
	return as.ToString(GetEnv(n, def))
}

//GetBool returns a environment variable as bool
func GetEnvBool(n string, def bool) bool {
	return as.ToBool(GetEnv(n, def))
}

//GetInt returns a environment variable as int
func GetEnvInt(n string, def int) int {
	return int(as.ToInt(GetEnv(n, def)))
}

//GetFloat returns a environment variable as float
func GetEnvFloat(n string, def float64) float64 {
	return as.ToFloat(GetEnv(n, def))
}

//prefixedName returns the prefixed name
func prefixedName(s string) string {
	if Prefix == "" {
		return s
	}
	return fmt.Sprintf("%s_%s", Prefix, s)
}
