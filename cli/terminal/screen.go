package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/robjporter/go-functions/as"
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

func GetTerminalSize() (int, int, error) {
	height := 0
	width := 0
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err == nil {
		splits := strings.Split(string(out), " ")
		if len(splits) == 2 {
			height = as.ToInt(splits[0])
			width = as.ToInt(strings.TrimRight(splits[1], "\n"))
		}
	}

	return height, width, err
}
