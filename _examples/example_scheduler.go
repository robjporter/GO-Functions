package main

import (
	"fmt"
	"runtime"
	"time"

	"../scheduler"
)

func main() {
	job := func() {
		t := time.Now()
		fmt.Println("JOB: >Time's up! @", t.UTC())
	}
	// Run every 2 seconds but not now.
	scheduler.Every(2).Seconds().NotImmediately().Run(job)

	// Run now and every X.
	scheduler.Every(1).Minutes().Run(job2)
	scheduler.Every().Day().Run(job)
	scheduler.Every().Monday().At("08:30").Run(job)

	// Keep the program from not exiting.
	runtime.Goexit()
}

func job2() {
	t := time.Now()
	fmt.Println("JOB 2: >Time's up! @", t.UTC())
}
