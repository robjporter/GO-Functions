package main

import (
	"fmt"
	"time"

	"../cache"
)

func main() {
	fmt.Println("START")
	Cache.Cache().Set("Hello", "HERE")
	Cache.Cache().SetKeyExpires("Hello", time.Second*2)
	fmt.Println(Cache.Cache().GetString("Hello"))
	time.Sleep(time.Second * 1)
	fmt.Println(Cache.Cache().GetString("Hello"))
	time.Sleep(time.Second * 2)
	fmt.Println(Cache.Cache().GetString("Hello"))
	fmt.Println("FINISH")
}
