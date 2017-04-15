package main

import (
	"fmt"

	"../tables"
)

func main() {
	fmt.Println("NORMAL TABLE  ********************************")
	table := tables.New([]string{"Name", "Host", "Type", "ID"})
	table.AddRow(map[string]interface{}{"Name": "MongoLab", "Host": "mongolab.com", "Type": "MongoDB Provider", "ID": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "Google App Engine", "Host": "appengine.google.com", "Type": "App Engine", "ID": "52518ff356357d17ec000004"})
	table.AddRow(map[string]interface{}{"Name": "Heroku", "Host": "heroku.com", "Type": "App Engine", "ID": "5251918e56357d17ec000005"})
	table.Print()

	fmt.Println("HORIZONTAL TABLE  ********************************")
	tables.PrintHorizontal(map[string]interface{}{"Name": "MongoLab", "Host": "mongolab.com", "Type": "MongoDB Provider", "ID": "52518c5d56357d17ec000002"})

	fmt.Println("MARKDOWN TABLE  ********************************")
	table3 := tables.New([]string{"Name", "Host", "Type", "ID"})
	table3.AddRow(map[string]interface{}{"Name": "MongoLab", "Host": "mongolab.com", "Type": "MongoDB Provider", "ID": "52518c5d56357d17ec000002"})
	table3.AddRow(map[string]interface{}{"Name": "Google App Engine", "Host": "appengine.google.com", "Type": "App Engine", "ID": "52518ff356357d17ec000004"})
	table3.AddRow(map[string]interface{}{"Name": "Heroku", "Host": "heroku.com", "Type": "App Engine", "ID": "5251918e56357d17ec000005"})
	table3.Markdown = true
	table3.Print()
}
