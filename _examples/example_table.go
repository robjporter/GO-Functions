package main

import (
    "../tables"
)

func main() {
	table := tables.NewTable([]string{"Name", "Host", "Type", "_id"})
	table.AddRow(map[string]interface{}{"Name": "MongoLab ", "Host": "mongolab.com", "Type": "MongoDB Provider", "_id": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "Google App Engine ", "Host": "appengine.google.com", "Type": "App Engine", "_id": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "Heroku ", "Host": "heroku.com", "Type": "App Engine", "_id": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "MySQL ", "Host": "mysql.com", "Type": "MySQL Provider", "_id": "52518c5d56357d17ec000002"})
	table.Print()

	tables.PrintHorizontal(map[string]interface{}{
		"Name": "Link",
		"Link": "mongolab.com",
	})

	table = tables.NewTable([]string{"Name", "Host", "Type", "_id"})
	table.AddRow(map[string]interface{}{"Name": "MongoLab ", "Host": "mongolab.com", "Type": "MongoDB Provider", "_id": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "Google App Engine ", "Host": "appengine.google.com", "Type": "App Engine", "_id": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "Heroku ", "Host": "heroku.com", "Type": "App Engine", "_id": "52518c5d56357d17ec000002"})
	table.AddRow(map[string]interface{}{"Name": "MySQL ", "Host": "mysql.com", "Type": "MySQL Provider", "_id": "52518c5d56357d17ec000002"})
	table.Markdown = true
	table.Print()
}