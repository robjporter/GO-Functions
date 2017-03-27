package main

import (
	"log"

	"../logs"
)

func main() {
	// create message zones and assign files for output
	newLog := map[string]interface{}{
		"error":   "error.log",
		"warning": "warning.log",
		"notice":  "notice.log",
	}

	// init logger
	logs.Init(newLog)

	// for ex. dump error
	logs.Write("error", "something had an error")

	// default output
	log.Println("Standard output to os.Stderr")

	// few more errors
	logs.Write("notice", "some notice")
	logs.Write("error", "another error")
	logs.Write("warning", "fancy warning")

	// wrong message zone, output to os.Stderr
	logs.Write("wrong_zone", "fail")
}
