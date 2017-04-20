package main

import (
	"fmt"

	"../browser"
)

func main() {
	/*
		ua1 := "(Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16"
		rs1 := browser.Parse(ua1)

		ua2 := "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.33 Safari/534.3 SE 2.X MetaSr 1.0"
		rs2 := browser.Parse(ua2)
	*/
	ua3 := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36"
	rs3 := browser.Parse(ua3)
	/* d
	fmt.Println("ua:", ua1,
		"\nbrowser:", rs1.Browser.Name,
		"\nversion:", rs1.Browser.Version,
		"\ndevice:", rs1.Device.Name,
		"\nos:", rs1.OS.Name,
		"\nversion:", rs1.OS.Version)

	fmt.Println()

	fmt.Println("ua:", ua2,
		"\nbrowser:", rs2.Browser.Name,
		"\nversion:", rs2.Browser.Version,
		"\ndevice:", rs2.Device.Name,
		"\nos:", rs2.OS.Name,
		"\nversion:", rs2.OS.Version)

	fmt.Println()
	*/
	fmt.Println("ua:", ua3,
		"\nbrowser:", rs3.Browser.Name,
		"\nversion:", rs3.Browser.Version,
		"\ndevice:", rs3.Device.Name,
		"\nos:", rs3.OS.Name,
		"\nversion:", rs3.OS.Version)
}
