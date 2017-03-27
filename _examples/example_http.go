package main

import (
	"fmt"

	"../http"
)

func main() {
	code, response, err := http.SendUnsecureHTTPSRequest("www.cisco.com", "get", "", nil)
	if err == nil {
		if code == 200 {
			fmt.Println(response)
		}
	}
}
