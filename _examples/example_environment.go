package main

import (
	"fmt"

	"../environment"
)

func main() {
	fmt.Println("COMPILED:                       >", environment.IsCompiled())
	fmt.Println("COMPILER:                       >", environment.Compiler())
	fmt.Println("ARCHITECTURE:                   >", environment.GOARCH())
	fmt.Println("CHECK ARCHITECTURE:             >", environment.CheckArchitecture())
	fmt.Println("BUILD STAMP:                    >", environment.BuildStamp())
	fmt.Println("BUILD HOST:                     >", environment.BuildHost())
	fmt.Println("BUILD USER:                     >", environment.BuildUser())
	fmt.Println("ENVIRONMENT DEBUG:              >", environment.BuildDebug())
	fmt.Println("ENVIRONMENT PATH SEPARATOR:     >", environment.PathSeparator())
	fmt.Println("ENVIRONMENT LIST SEPARATOR:     >", environment.ListSeparator())
	fmt.Println("GO OS:                          >", environment.GOOS())
	fmt.Println("GO ROOT:                        >", environment.GOROOT())
	fmt.Println("GO VERSION:                     >", environment.GOVER())
	fmt.Println("GO PATH:                        >", environment.GOPATH())
	fmt.Println("GO PATH BIN:                    >", environment.GOPATHBIN())
	fmt.Println("CPU:                            >", environment.NumCPU())
	fmt.Println("ENV STRING VALUE LOGNAME:       >", environment.GetEnvString("LOGNAME", ""))
	fmt.Println("ENV STRING VALUE BLOCKSIZE:     >", environment.GetEnvString("BLOCKSIZE", ""))
	fmt.Println("ENV STRING VALUE HOME:          >", environment.GetEnvString("HOME", ""))
	fmt.Println("ENV BOOL VALUE QT_HOMEBREW:     >", environment.GetEnvBool("QT_HOMEBREW", false))
	fmt.Println("ENV INT VALUE XPC_SERVICE_NAME: >", environment.GetEnvInt("XPC_SERVICE_NAME", 0))
	fmt.Println("ALL ENVIRONMENT VARIABLES:      >", environment.GetAllEnvironment())
}
