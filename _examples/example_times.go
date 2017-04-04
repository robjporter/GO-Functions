package main

import (
	"fmt"

	"../times"
)

func main() {
	t := times.New(2017, 4, 4, 12, 00, 00, "GMT")
	fmt.Println("ADD SIMPLE ===============================================")
	result, _ := t.AddYear()
	fmt.Println("ADD YEAR: ()          >", result)
	result, _ = t.AddMonth()
	fmt.Println("ADD MONTH: ()         >", result)
	result, _ = t.AddDay()
	fmt.Println("ADD DAY: ()           >", result)
	result, _ = t.AddHour()
	fmt.Println("ADD HOUR: ()          >", result)
	result, _ = t.AddMinute()
	fmt.Println("ADD MINUTE: ()        >", result)
	result, _ = t.AddSecond()
	fmt.Println("ADD SECOND: ()        >", result)

	fmt.Println("ADD ADVANCED =============================================")
	result, _ = t.AddYears(2)
	fmt.Println("ADD YEARS: (2)        >", result)
	result, _ = t.AddMonths(2)
	fmt.Println("ADD MONTHS: (2)       >", result)
	result, _ = t.AddDays(2)
	fmt.Println("ADD DAYS: (2)         >", result)
	result, _ = t.AddHours(2)
	fmt.Println("ADD HOURS: (2)        >", result)
	result, _ = t.AddMinutes(2)
	fmt.Println("ADD MINUTES: (2)      >", result)
	result, _ = t.AddSeconds(2)
	fmt.Println("ADD SECONDS: (2)      >", result)

	fmt.Println("SUB SIMPLE ===============================================")
	result, _ = t.SubYear()
	fmt.Println("SUB YEAR: ()          >", result)
	result, _ = t.SubMonth()
	fmt.Println("SUB MONTH: ()         >", result)
	result, _ = t.SubDay()
	fmt.Println("SUB DAY: ()           >", result)
	result, _ = t.SubHour()
	fmt.Println("SUB HOUR: ()          >", result)
	result, _ = t.SubMinute()
	fmt.Println("SUB MINUTE: ()        >", result)
	result, _ = t.SubSecond()
	fmt.Println("SUB SECOND: ()        >", result)

	fmt.Println("SUB ADVANCED =============================================")
	result, _ = t.SubYears(2)
	fmt.Println("SUB YEARS: (2)        >", result)
	result, _ = t.SubMonths(5)
	fmt.Println("SUB MONTHS: (5)       >", result)
	result, _ = t.SubDays(6)
	fmt.Println("SUB DAYS: (6)         >", result)
	result, _ = t.SubHours(2)
	fmt.Println("SUB HOURS: (2)        >", result)
	result, _ = t.SubMinutes(2)
	fmt.Println("SUB MINUTES: (2)      >", result)
	result, _ = t.SubSeconds(2)
	fmt.Println("SUB SECONDS: (2)      >", result)

	fmt.Println("FUNCTIONS ===============================================")
	fmt.Println("IS LEAP YEAR:         >", t.IsLeapYear())
}
