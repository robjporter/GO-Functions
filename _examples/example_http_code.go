package main

import (
	"fmt"
	"github.com/robjporter/go-functions/http"
)

func main() {
	fmt.Println(http.GetCodeShortInfo(100))
	fmt.Println(http.GetCodeLongInfo(100))
}
