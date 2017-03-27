package main

import (
	"fmt"

	"../network"
)

func main() {
	fmt.Println("Starting")
	ips := network.GetAllIP()
	fmt.Println("IPs found: ", len(ips))
	for i := 0; i < len(ips); i++ {
		fmt.Println(ips[i])
	}
	fmt.Println("===========================")
	ips2 := network.GetAllIPWithName()
	fmt.Println("Interfaces found: ", len(ips2))
	for k, v := range ips2 {
		fmt.Println(k, v)
	}
	fmt.Println("===========================")
	ips3 := network.GetMainIP()
	fmt.Println("Interfaces found: ", len(ips3))
	for i := 0; i < len(ips3); i++ {
		fmt.Println(ips3[i])
	}
}
