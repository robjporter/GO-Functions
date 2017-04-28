package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/robjporter/go-functions/astilectron"
)

func main() {
	// Parse flags
	flag.Parse()

	// Start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <title>Hello world</title>
		</head>
		<body>
		    Hello world
		</body>
		</html>`))
	})
	go http.ListenAndServe("127.0.0.1:4000", nil)

	// Create astilectron
	var a *astilectron.Astilectron
	var err error
	if a, err = astilectron.New(astilectron.Options{BaseDirectoryPath: os.Getenv("GOPATH") + "/src/github.com/asticode/go-astilectron/examples"}); err != nil {
		fmt.Println("Creating new astilectron failed", err)
	}
	defer a.Close()
	a.HandleSignals()
	a.On(astilectron.EventNameAppStop, func(e astilectron.Event) (deleteListener bool) {
		a.Stop()
		return
	})

	// Start
	if err = a.Start(); err != nil {
		fmt.Println("Starting failed: ", err)
	}

	// Create window
	var w *astilectron.Window
	if w, err = a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(1000),
	}); err != nil {
		fmt.Println("New window failed: ", err)
	}
	if err = w.Create(); err != nil {
		fmt.Println("Creating window failed: ", err)
	}

	// Blocking pattern
	a.Wait()
}
