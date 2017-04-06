package main

import (
	"fmt"

	"time"

	"../times"
)

func main() {
	ti := time.Now()
	year := ti.Year()
	month := ti.Month()
	day := ti.Day()
	hour := ti.Hour()
	minute := ti.Minute()
	second := ti.Second()

	//t := times.New(year, times.MonthNameToNumber(month), day, hour, minute, second, "Europe/London")
	//t := times.Today("Europe/London")
	t := times.TodayAuto()
	t6 := t.Copy()
	fmt.Println(t6)
	fmt.Println("ADD SIMPLE ===============================================")
	result, _ := t.AddDecade()
	fmt.Println("ADD DECADE:               ()>", result)
	result, _ = t.AddYear()
	fmt.Println("ADD YEAR:                 ()>", result)
	result, _ = t.AddMonth()
	fmt.Println("ADD MONTH:                ()>", result)
	result, _ = t.AddWeek()
	fmt.Println("ADD WEEK:                 ()>", result)
	result, _ = t.AddDay()
	fmt.Println("ADD DAY:                  ()>", result)
	result, _ = t.AddHour()
	fmt.Println("ADD HOUR:                 ()>", result)
	result, _ = t.AddMinute()
	fmt.Println("ADD MINUTE:               ()>", result)
	result, _ = t.AddSecond()
	fmt.Println("ADD SECOND:               ()>", result)

	fmt.Println("\nADD ADVANCED =============================================")
	result, _ = t.AddDecades(2)
	fmt.Println("ADD DECADES:             (2)>", result)
	result, _ = t.AddYears(2)
	fmt.Println("ADD YEARS:               (2)>", result)
	result, _ = t.AddMonths(2)
	fmt.Println("ADD MONTHS:              (2)>", result)
	result, _ = t.AddWeeks(2)
	fmt.Println("ADD WEEKS:               (2)>", result)
	result, _ = t.AddDays(2)
	fmt.Println("ADD DAYS:                (2)>", result)
	result, _ = t.AddHours(2)
	fmt.Println("ADD HOURS:               (2)>", result)
	result, _ = t.AddMinutes(2)
	fmt.Println("ADD MINUTES:             (2)>", result)
	result, _ = t.AddSeconds(2)
	fmt.Println("ADD SECONDS:             (2)>", result)

	fmt.Println("\nSUB SIMPLE ===============================================")
	result, _ = t.SubDecade()
	fmt.Println("SUB DECADE:              ()>", result)
	result, _ = t.SubYear()
	fmt.Println("SUB YEAR:                ()>", result)
	result, _ = t.SubMonth()
	fmt.Println("SUB MONTH:               ()>", result)
	result, _ = t.SubWeek()
	fmt.Println("SUB WEEK:                ()>", result)
	result, _ = t.SubDay()
	fmt.Println("SUB DAY:                 ()>", result)
	result, _ = t.SubHour()
	fmt.Println("SUB HOUR:                ()>", result)
	result, _ = t.SubMinute()
	fmt.Println("SUB MINUTE:              ()>", result)
	result, _ = t.SubSecond()
	fmt.Println("SUB SECOND:              ()>", result)

	fmt.Println("\nSUB ADVANCED =============================================")
	result, _ = t.SubDecades(2)
	fmt.Println("SUB DECADES:            (2)>", result)
	result, _ = t.SubYears(2)
	fmt.Println("SUB YEARS:              (2)>", result)
	result, _ = t.SubMonths(5)
	fmt.Println("SUB MONTHS:             (5)>", result)
	result, _ = t.SubWeeks(2)
	fmt.Println("SUB WEEKS:              (2)>", result)
	result, _ = t.SubDays(6)
	fmt.Println("SUB DAYS:               (6)>", result)
	result, _ = t.SubHours(2)
	fmt.Println("SUB HOURS:              (2)>", result)
	result, _ = t.SubMinutes(2)
	fmt.Println("SUB MINUTES:            (2)>", result)
	result, _ = t.SubSeconds(2)
	fmt.Println("SUB SECONDS:            (2)>", result)

	fmt.Println("\nSTART ===================================================")
	result, _ = t.StartOfCentury()
	fmt.Println("START OF CENTURY:       ()>", result)
	result, _ = t.StartOfDecade()
	fmt.Println("START OF DECADE:        ()>", result)
	result, _ = t.StartOfYear()
	fmt.Println("START OF YEAR:          ()>", result)
	result, _ = t.StartOfMonth()
	fmt.Println("START OF MONTH:         ()>", result)
	result, _ = t.StartOfWorkWeek()
	fmt.Println("START OF WORK WEEK:     ()>", result)
	result, _ = t.StartOfWeek()
	fmt.Println("START OF WEEK:          ()>", result)
	result, _ = t.StartOfDay()
	fmt.Println("START OF DAY:           ()>", result)
	result, _ = t.StartOfHour()
	fmt.Println("START OF HOUR:          ()>", result)
	result, _ = t.StartOfMinute()
	fmt.Println("START OF MINUTE:        ()>", result)

	fmt.Println("\nEND =====================================================")
	result, _ = t.EndOfCentury()
	fmt.Println("END OF CENTURY:         ()>", result)
	result, _ = t.EndOfDecade()
	fmt.Println("END OF DECADE:          ()>", result)
	result, _ = t.EndOfYear()
	fmt.Println("END OF YEAR:            ()>", result)
	result, _ = t.EndOfMonth()
	fmt.Println("END OF MONTH:           ()>", result)
	result, _ = t.EndOfWorkWeek()
	fmt.Println("END OF WORK WEEK:       ()>", result)
	result, _ = t.EndOfWeek()
	fmt.Println("END OF WEEK:            ()>", result)
	result, _ = t.EndOfDay()
	fmt.Println("END OF DAY:             ()>", result)
	result, _ = t.EndOfHour()
	fmt.Println("END OF HOUR:            ()>", result)
	result, _ = t.EndOfMinute()
	fmt.Println("END OF MINUTE:          ()>", result)

	fmt.Println("\nQUARTER  ================================================")
	fmt.Println("Is 1st QUARTER:         ()>", t.Is1stQuarter())
	fmt.Println("Is 2nd QUARTER:         ()>", t.Is2ndQuarter())
	fmt.Println("Is 3rd QUARTER:         ()>", t.Is3rdQuarter())
	fmt.Println("Is 4th QUARTER:         ()>", t.Is4thQuarter())
	result, _ = t.StartOfQuarter()
	fmt.Println("START OF QUARTER:       ()>", result)
	result, _ = t.EndOfQuarter()
	fmt.Println("END OF QUARTER:         ()>", result)
	fmt.Println("QUARTER:                ()>", t.Quarter())
	fmt.Println("QUARTER NUMBER:         ()>", t.QuarterNumber())

	fmt.Println("\nSEASONS =================================================")
	fmt.Println("SEASON:                 ()>", t.Season())
	fmt.Println("TIME TO SPRING:         ()>", t.TimeToSpring())
	fmt.Println("TIME TO SPRING:         ()>", t.TimeToSpringDiff())
	fmt.Println("TIME TO SUMMER:         ()>", t.TimeToSummer())
	fmt.Println("TIME TO SUMMER:         ()>", t.TimeToSummerDiff())
	fmt.Println("TIME TO AUTUMN:         ()>", t.TimeToAutumn())
	fmt.Println("TIME TO AUTUMN:         ()>", t.TimeToAutumnDiff())
	fmt.Println("TIME TO WINTER:         ()>", t.TimeToWinter())
	fmt.Println("TIME TO WINTER:         ()>", t.TimeToWinterDiff())

	fmt.Println("\nTAX YEAR ================================================")
	fmt.Println("TAX YEAR:               ()>", t.TaxYear())
	result, _ = t.StartOfTaxYear()
	fmt.Println("START OF UK TAX YEAR:   ()>", result)
	result, _ = t.EndOfTaxYear()
	fmt.Println("END OF UK TAX YEAR:     ()>", result)
	fmt.Println("NEXT UK TAX YEAR:       ()>", t.TimeToTaxYear())
	fmt.Println("NEXT UK TAX YEAR DIFF   ()>", t.TimeToTaxYearDiff())

	fmt.Println("\nIS ======================================================")
	fmt.Println("IS WEEKEND:             ()>", t.IsWeekend())
	fmt.Println("IS WORKDAY:             ()>", t.IsWorkday())
	fmt.Println("IS MONDAY:              ()>", t.IsMonday())
	fmt.Println("IS TUESDAY:             ()>", t.IsTuesday())
	fmt.Println("IS WEDNESDAY:           ()>", t.IsWednesday())
	fmt.Println("IS THURSDAY:            ()>", t.IsThursday())
	fmt.Println("IS FRIDAY:              ()>", t.IsFriday())
	fmt.Println("IS SATURDAY:            ()>", t.IsSaturday())
	fmt.Println("IS SUNDAY:              ()>", t.IsSunday())
	t4 := times.New(year-1, times.MonthNameToNumber(month), day, hour, minute, second, "Europe/London")
	t5 := times.New(year+1, times.MonthNameToNumber(month), day, hour, minute, second, "Europe/London")
	fmt.Println("IS BETWEEN:        (t4,t5)>", t.IsBetween(t4, t5))
	fmt.Println("IS SPRING:              ()>", t.IsSpring())
	fmt.Println("IS SUMMER:              ()>", t.IsSummer())
	fmt.Println("IS AUTUMN:              ()>", t.IsAutumn())
	fmt.Println("IS WINTER:              ()>", t.IsWinter())

	fmt.Println("\nDIFFERENCE ==============================================")
	t3 := times.New(year+1, times.MonthNameToNumber(month)+1, day+8, hour+1, minute+1, second+4, "Europe/London")
	fmt.Println("DIFFERENCE:             ()>", t.DifferenceDiff(t3))
	fmt.Println("DIFFERENCE:             ()>", t.Difference(t3))
	fmt.Println("DIFFINYEARS:            ()>", t.DiffInYears(t3))
	fmt.Println("DIFFINMONTHS:           ()>", t.DiffInMonths(t3))
	fmt.Println("DIFFINWEEKS:            ()>", t.DiffInWeeks(t3))
	fmt.Println("DIFFINDAYS:             ()>", t.DiffInDays(t3))
	fmt.Println("DIFFINHOURS:            ()>", t.DiffInHours(t3))
	fmt.Println("DIFFINMINUTES:          ()>", t.DiffInMinutes(t3))
	fmt.Println("DIFFINSECONDS:          ()>", t.DiffInSeconds(t3))

	fmt.Println("\nFORMATS =================================================")
	result2, _ := t.Format822()
	fmt.Println("FORMAT 822:             ()>", result2)
	result2, _ = t.Format1123()
	fmt.Println("FORMAT 1123:            ()>", result2)
	result2, _ = t.Format1123z()
	fmt.Println("FORMAT 1123z:           ()>", result2)
	result2, _ = t.Format3339()
	fmt.Println("FORMAT 3339:            ()>", result2)
	result2, _ = t.Format3339nano()
	fmt.Println("FORMAT 3339 Nano:       ()>", result2)
	result2, _ = t.Format8222z()
	fmt.Println("FORMAT 8222z:           ()>", result2)
	result2, _ = t.Format850()
	fmt.Println("FORMAT 850:             ()>", result2)
	result2, _ = t.Format1()
	fmt.Println("FORMAT 1:               ()>", result2)
	result2, _ = t.Format2()
	fmt.Println("FORMAT 2:               ()>", result2)
	result2, _ = t.Format3()
	fmt.Println("FORMAT 3:               ()>", result2)
	result2, _ = t.Format4()
	fmt.Println("FORMAT 4:               ()>", result2)
	result2, _ = t.Format5()
	fmt.Println("FORMAT 5:               ()>", result2)
	result2, _ = t.Format6()
	fmt.Println("FORMAT 6:               ()>", result2)
	result2, _ = t.Format7()
	fmt.Println("FORMAT 7:               ()>", result2)
	result2, _ = t.Format8()
	fmt.Println("FORMAT 8:               ()>", result2)
	result2, _ = t.Format9()
	fmt.Println("FORMAT 9:               ()>", result2)
	result2, _ = t.Format10()
	fmt.Println("FORMAT 10:              ()>", result2)
	result2, _ = t.Format11()
	fmt.Println("FORMAT 11:              ()>", result2)
	result2, _ = t.Format12()
	fmt.Println("FORMAT 12:              ()>", result2)
	result2, _ = t.Format13()
	fmt.Println("FORMAT 13:              ()>", result2)
	result2, _ = t.Format14()
	fmt.Println("FORMAT 14:              ()>", result2)
	result2, _ = t.Format15()
	fmt.Println("FORMAT 15:              ()>", result2)
	result2, _ = t.Format16()
	fmt.Println("FORMAT 16:              ()>", result2)
	result2, _ = t.Format17()
	fmt.Println("FORMAT 17:              ()>", result2)
	result2, _ = t.Format18()
	fmt.Println("FORMAT 18:              ()>", result2)
	result2, _ = t.Format19()
	fmt.Println("FORMAT 19:              ()>", result2)
	result2, _ = t.Format20()
	fmt.Println("FORMAT 20:              ()>", result2)

	fmt.Println("\nFUNCTIONS ===============================================")
	t = times.New(year, times.MonthNameToNumber(month), day, hour, minute, second, "Europe/London")
	fmt.Println("IS LEAP YEAR:           ()>", t.IsLeapYear())
	fmt.Println("NEXT LEAP YEAR:         ()>", t.NextLeapYear())
	result, _ = t.TimeNext(time.Wednesday)
	fmt.Println("TIME NEXT:     (Wednesday)>", result)
	result, _ = t.TimePrevious(time.Wednesday)
	fmt.Println("TIME LAST:     (Wednesday)>", result)

	t2 := times.New(year, times.MonthNameToNumber(month), day+1, hour, minute, second, "Europe/London")
	fmt.Println("IS FUTURE:              ()>", t.IsFuture(t2))
	t2.SubDay()
	fmt.Println("IS PAST:                ()>", t.IsPast(t2))

	result2, _ = t.ISOWeek()
	fmt.Println("ISO WEEK:               ()>", result2)
	fmt.Println("LOCATION:               ()>", t.Location)

}
