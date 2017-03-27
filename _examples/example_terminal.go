package main

import (
	"fmt"

	"../terminal"
)

func main() {
	terminal.ClearScreen()
	//fmt.Println(terminal.Width(), " x ", terminal.Height())
	terminal.WaitD(4)
	terminal.ClearLine()
	terminal.Wait(4, 1)
	terminal.PrintXY(10, 10, "HERE\n")
	//cx, cy := terminal.CursorXY()
	//fmt.Println(cx, cy)
	//
	fmt.Println("Questions")
	answer := terminal.AskString("This is a test question")
	fmt.Println("ANSWER:> ", answer)
	answer2 := terminal.YesOrNo("This is a test?", false)
	fmt.Println("ANSWER:> ", answer2)
	answer3 := terminal.Choice("This is a test", []string{"Answer1", "Answer2"})
	fmt.Println("ANSWER:> ", answer3)
}
