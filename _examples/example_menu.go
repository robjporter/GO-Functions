package main

import (
	"fmt"
	"os"

	climenu "../menu"
)

func callback(id string) {
	fmt.Println("Chose item:", id)
}

func main() {
	menu2 := climenu.NewMenu("", "Choose an action", 0)
	menu2.AddMenuItem("Sub Menu1", "connection")
	menu2.AddMenuItem("Sub Menu1", "tenants")
	menu2.AddMenuItem("Back", "back")

	menu3 := climenu.NewMenu("", "Choose an action", 0)
	menu3.AddMenuItem("Sub Menu2", "connection")
	menu3.AddMenuItem("Sub Menu2", "tenants")
	menu3.AddMenuItem("Back", "back")

	menu := climenu.NewButtonMenu("", "Choose an action")
	menu.AddMenuItem("Setup ACI System Connection", "connection").SetSubMenu(menu2)
	menu.AddMenuItem("Manage Tenants", "tenants").SetSubMenu(menu3)
	menu.AddMenuItem("Quit", "quit")
	action, escaped := menu.Run()
	if escaped || action == "quit" {
		os.Exit(0)
	}

}
