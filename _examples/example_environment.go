package main

import (
	"../environment"
	"fmt"
)

func main() {
	fmt.Println("COMPILED: ", environment.IsCompiled())
	fmt.Println("COMPILER: ", environment.Compiler())
	fmt.Println("Architecture: ", environment.GOARCH())
	fmt.Println("GO OS: ", environment.GOOS())
	fmt.Println("GO Root: ", environment.GOROOT())
	fmt.Println("GO Version: ", environment.GOVER())
	fmt.Println("GO Path: ", environment.GOPATH())
	fmt.Println("CPU: ", environment.NumCPU())
	fmt.Println("ENV STRING VALUE LOGNAME: ", environment.GetEnvString("LOGNAME", ""))
	fmt.Println("ENV STRING VALUE BLOCKSIZE: ", environment.GetEnvString("BLOCKSIZE", ""))
	fmt.Println("ENV STRING VALUE HOME: ", environment.GetEnvString("HOME", ""))
	fmt.Println("ENV BOOL VALUE QT_HOMEBREW: ", environment.GetEnvBool("QT_HOMEBREW", false))
	fmt.Println("ENV INT VALUE XPC_SERVICE_NAME: ", environment.GetEnvInt("XPC_SERVICE_NAME", 0))
}
