package terminal

import (
	"fmt"
	"strconv"
	"strings"
)

const ESC = "\033["

func clearscreen(i int) {
	fmt.Printf(ESC+"%dJ", i)
}

func clearline(i int) {
	fmt.Printf(ESC+"%dK", i)
}

func ClearLine() {
	clearline(2)
}

func ClearScr() {
	clearscreen(2)
	GotoXY(1, 1)
}

func GotoXY(x int, y int) {
	fmt.Printf(ESC+"%d;%dH", x, y)
}
func CursorUp(i int) {
	fmt.Printf(ESC+"%dA", i)
}

func CursorDn(i int) {
	fmt.Printf(ESC+"%dB", i)
}

func CursorRt(i int) {
	fmt.Printf(ESC+"%dC", i)
}

func CursorLf(i int) {
	fmt.Printf(ESC+"%dD", i)
}

func CursorSave() {
	fmt.Print(ESC + "s")
}

func CursorRestore() {
	fmt.Print(ESC + "u")
}

func PrintXY(x int, y int, s string) {
	GotoXY(x, y)
	fmt.Printf("%s", s)
}

func CursorXY() (int, int) {
	var ch int
	var line string
	var ret []int

	fmt.Print(ESC + "6n")

	for ch != 'R' {
		ch = RawKey()
		if ch != 27 && ch != '[' {
			ret = append(ret, ch)
		}
	}

	for _, value := range ret {
		line += string(int(value))
	}

	line = strings.TrimRight(line, "R")
	parts := strings.Split(line, ";")

	X, err := strconv.Atoi(parts[0])
	if err == nil {

	}

	Y, err := strconv.Atoi(parts[1])
	if err == nil {

	}

	return X, Y
}
