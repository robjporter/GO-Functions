package main

import (
	"fmt"

	"./config"
)

func main() {
	cfg := config.New()
	cfg.ReadFiles("./config.yaml", "./conf.json")
	data := `{"menu": {"id": "file","value": "File","popup": {"menuitem": [{"value": "New", "onclick": "CreateNewDoc()"},{"value": "Open", "onclick": "OpenDoc()"},{"value": "Close", "onclick": "CloseDoc()"}]}}}`
	cfg.ReadString(data)
	//cfg.SaveYaml("test.yaml")
	//fmt.Println("FILES: >", cfg.GetConfigFiles())
	//fmt.Println(cfg.AllSettings())
	fmt.Println("VALUE: >", cfg.Get("menu.popup.menuitem.0.onclick"))
	fmt.Println("VALUE: >", cfg.Get("ucs.systems.0.url"))

	//fmt.Println(cfg.GetString("menu.id"))
	//fmt.Println(cfg.GetString("output.file"))
	//fmt.Println("GET************************************")
	//fmt.Println("UCS: >", cfg.Get("ucs"))
	//fmt.Println("")
	//fmt.Println("SYSTEMS: >", cfg.Get("ucs.systems"))
	//fmt.Println("")
	//fmt.Println("ELEMENT: >", cfg.Get("ucs.systems.0"))
	//fmt.Println("")
	//fmt.Println("ELEMENT: >", cfg.GetString("ucs.systems.0.url"))
}
