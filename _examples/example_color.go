package main

import (
	"fmt"

	"../colors"
)

func main() {
	fmt.Println("TEXT EFFECTS**********************************************")
	fmt.Println(colors.Bold("BOLD"))
	fmt.Println(colors.Italic("ITALIC"))
	fmt.Println(colors.Underline("UNDERLINE"))
	fmt.Println(colors.StrikeThrough("STRIKETHROUGH"))
	fmt.Println(colors.Blink("BLINK"))
	fmt.Println(colors.Reversed("REVERSED"))

	fmt.Println("\nTEXT COLORS**********************************************")
	fmt.Println(colors.Background("BACKGROUND", colors.RED))
	fmt.Println(colors.Color("COLOR", colors.YELLOW))
	fmt.Println(colors.Highlight("HIGHLIGHT", "GHLI", colors.BRIGHTYELLOW))
	fmt.Println(colors.Color(colors.Bold("BOLD YELLOW"), colors.YELLOW))
	fmt.Println(colors.Background(colors.Color(colors.Bold("BOLD YELLOW WITH RED BACKGROUND"), colors.YELLOW), colors.RED))
	fmt.Println(colors.Color("BLACK", colors.BLACK))
	fmt.Println(colors.Color("BRIGHTBLACK", colors.BRIGHTBLACK))
	fmt.Println(colors.Color("RED", colors.RED))
	fmt.Println(colors.Color("BRIGHTRED", colors.BRIGHTRED))
	fmt.Println(colors.Color("GREEN", colors.GREEN))
	fmt.Println(colors.Color("BRIGHTGREEN", colors.BRIGHTGREEN))
	fmt.Println(colors.Color("YELLOW", colors.YELLOW))
	fmt.Println(colors.Color("BRIGHTYELLOW", colors.BRIGHTYELLOW))
	fmt.Println(colors.Color("BLUE", colors.BLUE))
	fmt.Println(colors.Color("BRIGHTBLUE", colors.BRIGHTBLUE))
	fmt.Println(colors.Color("MAGENTA", colors.MAGENTA))
	fmt.Println(colors.Color("BRIGHTMAGENTA", colors.BRIGHTMAGENTA))
	fmt.Println(colors.Color("WHITE", colors.WHITE))
	fmt.Println(colors.Color("BRIGHTWHITE", colors.BRIGHTWHITE))

	fmt.Println("\nTEXT PANELS********************************************")
	fmt.Println(colors.BlackSmallPanel("Here is some text in a black panel."))
	fmt.Println(colors.RedSmallPanel("Here is some text in a red panel."))
	fmt.Println(colors.GreenSmallPanel("Here is some text in a green panel."))
	fmt.Println(colors.YellowSmallPanel("Here is some text in a yellow panel."))
	fmt.Println(colors.BlueSmallPanel("Here is some text in a blue panel."))
	fmt.Println(colors.MagentaSmallPanel("Here is some text in a magenta panel."))
	fmt.Println(colors.CyanSmallPanel("Here is some text in a cyan panel."))
	fmt.Println(colors.WhiteSmallPanel("Here is some text in a white panel."))
	fmt.Println(colors.BlackPanel("Here is some text in a black panel."))
	fmt.Println(colors.RedPanel("Here is some text in a red panel."))
	fmt.Println(colors.GreenPanel("Here is some text in a green panel."))
	fmt.Println(colors.YellowPanel("Here is some text in a yellow panel."))
	fmt.Println(colors.BluePanel("Here is some text in a blue panel."))
	fmt.Println(colors.MagentaPanel("Here is some text in a magenta panel."))
	fmt.Println(colors.CyanPanel("Here is some text in a cyan panel."))
	fmt.Println(colors.WhitePanel("Here is some text in a white panel."))

	fmt.Println("\nTEXT COMMANDS********************************************")
	fmt.Println(colors.Title("Title"))
	fmt.Println(colors.CustomTitle("Title", colors.BRIGHTWHITE, colors.BRIGHTYELLOW))
	fmt.Println(colors.Info("This is an info message"))
	fmt.Println(colors.Success("This is a success message"))
	fmt.Println(colors.Warning("This is an warning message"))
	fmt.Println(colors.Error("This is an error message"))

	fmt.Println("\nASCII ICONS**********************************************")
	fmt.Println(colors.TICK, colors.Green(colors.TICK), colors.BrightGreen(colors.TICK))
	fmt.Println(colors.CROSS, colors.Red(colors.CROSS), colors.BrightRed(colors.CROSS))
	fmt.Println(colors.COPYRIGHT)
	fmt.Println(colors.REGISTREDTM)
	fmt.Println(colors.TRADEMARK)
	fmt.Println(colors.BULLET)
	fmt.Println(colors.ARROWLEFT)
	fmt.Println(colors.ARROWRIGHT)
	fmt.Println(colors.ARROWUP)
	fmt.Println(colors.ARROWDOWN)
	fmt.Println(colors.ARROWLEFTRIGHT)
	fmt.Println(colors.INFINITY)
	fmt.Println(colors.CELSIUS)
	fmt.Println(colors.FAHRENHEIT)
	fmt.Println(colors.SUNSHINE)
	fmt.Println(colors.CLOUDY)
	fmt.Println(colors.RAIN)
	fmt.Println(colors.SNOW)
	fmt.Println(colors.STARBLACK)
	fmt.Println(colors.STARWHITE)
	fmt.Println(colors.PHONEBLACK)
	fmt.Println(colors.PHONEWHITE)
	fmt.Println(colors.POINTLEFT)
	fmt.Println(colors.POINTRIGHT)
	fmt.Println(colors.POINTUP)
	fmt.Println(colors.POINTDOWN)
	fmt.Println(colors.DEATH)
	fmt.Println(colors.SMILEY)
	fmt.Println(colors.HEART)
	fmt.Println(colors.DIAMOND)
	fmt.Println(colors.SPADE)
	fmt.Println(colors.CLUB)

	colors.PrintQColor(7, 1, "TESTING STRING")
}
