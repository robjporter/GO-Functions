package main

import (
	"fmt"

	"../regex"
)

func main() {
	text := "This is a test and my ip is 192.168.1.1/24, my email address is test@test.com and test@gmail.com, "
	text += "out IPv6 address is 3ffe:1900:4545:3:200:f8ff:fe21:67cf and also fe80::200:f8ff:fe21:67cf"
	text += "with the web address of http://www.test.com and the secure address https://www.test.com"
	text += ", this site contains HTML, which looks like <b>html</b>."
	text += "The web server has a mac address of 00:11:22:33:44:55, it costs $22.22 a month, which is about £18."
	text += "This service started on the 28/04/2008 @ 12:00 or 13:30 and is located at SS2 2SS."
	text += "this is for the custom test: abb and abbbb"
	text += " 2.2(8g)  3.1(2f)  2.1(3l)  3.0(2f)"

	password1 := "A2345B"
	password2 := "A2345!CD"
	password3 := "A2345B"
	ucs := "2.2(8g)"

	fmt.Println("IP v4:          >", regex.IP(text))
	fmt.Println("IP v4 CIDR:     >", regex.IPCidr(text))
	fmt.Println("IP v6:          >", regex.IPv6(text))
	fmt.Println("EMAIL:          >", regex.Email(text))
	fmt.Println("FREE EMAIL:     >", regex.EmailFree(text))
	fmt.Println("HTTP:           >", regex.Http(text))
	fmt.Println("HTTPS:          >", regex.Https(text))
	fmt.Println("URLs:           >", regex.URL(text))
	fmt.Println("MAC:            >", regex.Mac(text))
	fmt.Println("DOLLAR:         >", regex.Dollar(text))
	fmt.Println("POUND:          >", regex.Pound(text))
	fmt.Println("DATE:           >", regex.Date(text))
	fmt.Println("TIME:           >", regex.Time24(text))
	fmt.Println("UK POSTCODE:    >", regex.PostcodeUK(text))
	fmt.Println("PASSWORD LOW:   >", regex.PasswordLow(password1))
	fmt.Println("PASSWORD MED:   >", regex.PasswordMedium(password2))
	fmt.Println("PASSWORD HIG:   >", regex.PasswordHigh(password3))
	fmt.Println("UCS VERSION:    >", regex.UCSVersion(text))
	fmt.Println("UCS VERSION:    >", regex.UCSVersion(ucs))
	fmt.Println("CLEAN STRING:   >", regex.CleanString(password2))

	name := "test"
	custom := "ab{2,}"
	regex.AddPattern(name, custom)
	fmt.Println("CUSTOM TEST:    >", regex.Pattern(name, text))
}
