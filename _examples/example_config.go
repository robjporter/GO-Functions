package main

import (
	"fmt"

	"../config"
)

func main() {
	cfg := config.New()
	//cfg.ReadFiles("./config.yaml", "./conf.json")
	data := `{"menu": {"id": "file","value": "File","popup": {"menuitem": [{"value": "New", "onclick": "CreateNewDoc()"},{"value": "Open", "onclick": "OpenDoc()"},{"value": "Close", "onclick": "CloseDoc()"}]}}}`
	cfg.ReadString(data)

	if size, _ := cfg.GetSliceSize("menu.popup.menuitem"); size > 0 {
		tmp2, err := cfg.GetSlice("menu.popup.menuitem")
		if err == nil {
			for i := 0; i < len(tmp2); i++ {
				fmt.Println(tmp2[i])
				tmp3, _ := cfg.GetInterfaceMapSliceElement(tmp2[i])
				fmt.Println(tmp3["value"])
			}
		}
	}
}
