package main

import (
	"fmt"
	"github.com/robjporter/go-functions/cisco/ucs"
)

func main() {
	myucs := ucs.New()

	fmt.Println("SECTION 1 •******************************")
	myucs.Login("10.52.208.160", "admin", "C1sco123")
	if myucs.LastResponse.Errors == nil {
		fmt.Println("VERSION:     >", myucs.GetVersion())
		fmt.Println("PRIVILEDGES: >", myucs.GetPriviledges())
		fmt.Println("IS ADMIN:    >", myucs.IsAdmin())
		//fmt.Println(myucs.LastResponse.Response)
		//fmt.Println(myucs.LastResponse.Body)
	}
	myucs.Logout()

	fmt.Println("\nSECTION 2 ************************")
	resp2, err2 := ucs.New().Login("10.52.208.160", "admin", "C1sco123").End()
	fmt.Println(resp2)
	fmt.Println(err2)

	fmt.Println("\nSECTION 3 ************************")
	myucs3 := ucs.New()
	myucs3.Login("10.52.208.160", "admin", "C1sco123")
	fmt.Println(myucs3.GetVersion())
	myucs3.Logout()
}
