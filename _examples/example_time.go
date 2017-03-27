package main

import (
	"fmt"
	"time"

	"../time"
)

func main() {
	timing.Timer("TEST")
	time.Sleep(2 * time.Second)
	timing.Timer("TEST2")
	time.Sleep(2 * time.Second)
	fmt.Println("TEST:> ", timing.Timer("TEST"))
	fmt.Println("TEST2:> ", timing.Timer("TEST2"))
}
