package terminal

import (
	"fmt"
	"time"
)

var (
	styles = map[int][]string{
		0: {"\010|", "\010/", "\010-", "\010\\", "\010"},
		1: {"\010←", "\010↖", "\010↑", "\010↗", "\010→", "\010↘", "\010↓", "\010↙"},
	}
)

func WaitD(i int) {
	Wait(i, 0)
}
func Wait(i int, style int) {
	for x := 1; x < i*5; x++ {
		for y := 0; y < len(styles[style]); y++ {
			fmt.Print(styles[style][y])
			time.Sleep(50 * time.Millisecond)
		}
		/*
			fmt.Print("|")
			time.Sleep(50 * time.Millisecond)
			fmt.Print("\010/")
			time.Sleep(50 * time.Millisecond)
			fmt.Print("\010-")
			time.Sleep(50 * time.Millisecond)
			fmt.Print("\010\\")
			time.Sleep(50 * time.Millisecond)
			fmt.Print("\010")
		*/
	}
}
