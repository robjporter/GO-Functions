package main

import (
	"fmt"
	"time"

	"../now"
)

func main() {
	time.Now() // 2013-11-18 17:51:49.123456789 Mon

	fmt.Println("Beginning of Minute: ", now.BeginningOfMinute()) // 2013-11-18 17:51:00 Mon
	fmt.Println("Beginning of Hour: ", now.BeginningOfHour())     // 2013-11-18 17:00:00 Mon
	fmt.Println("Beginning of Day: ", now.BeginningOfDay())       // 2013-11-18 00:00:00 Mon
	fmt.Println("Beginning of Week: ", now.BeginningOfWeek())     // 2013-11-17 00:00:00 Sun

	now.FirstDayMonday = true // Set Monday as first day, default is Sunday
	fmt.Println("Change first day of week to Monday=============================")
	fmt.Println("Beginning of Minute: ", now.BeginningOfMinute()) // 2013-11-18 00:00:00 Mon
	fmt.Println("Beginning of Hour: ", now.BeginningOfHour())     // 2013-11-01 00:00:00 Fri
	fmt.Println("Beginning of Day: ", now.BeginningOfDay())       // 2013-10-01 00:00:00 Tue
	fmt.Println("Beginning of Week: ", now.BeginningOfWeek())     // 2013-01-01 00:00:00 Tue

	fmt.Println("Time===========================================================")
	fmt.Println("End of Minute: ", now.EndOfMinute()) // 2013-11-18 17:51:59.999999999 Mon
	fmt.Println("End of Hour: ", now.EndOfHour())     // 2013-11-18 17:59:59.999999999 Mon
	fmt.Println("End of Day: ", now.EndOfDay())       // 2013-11-18 23:59:59.999999999 Mon
	fmt.Println("End of Week: ", now.EndOfWeek())     // 2013-11-23 23:59:59.999999999 Sat

	fmt.Println("Change first day of week to Monday=============================")
	now.FirstDayMonday = true                           // Set Monday as first day, default is Sunday
	fmt.Println("End of Week: ", now.EndOfWeek())       // 2013-11-24 23:59:59.999999999 Sun
	fmt.Println("End of Month: ", now.EndOfMonth())     // 2013-11-30 23:59:59.999999999 Sat
	fmt.Println("End of Quarter: ", now.EndOfQuarter()) // 2013-12-31 23:59:59.999999999 Tue
	fmt.Println("End of Year: ", now.EndOfYear())       // 2013-12-31 23:59:59.999999999 Tue

	// Use another time
	t := time.Date(2012, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
	fmt.Println("Change Time====================================================")
	fmt.Println("End of Month: ", now.New(t).EndOfMonth()) // 2013-02-28 23:59:59.999999999 Thu

	// Don't want be bothered with the First Day setting, Use Monday, Sunday
	fmt.Println("Get Specific days==============================================")
	fmt.Println("Mondy: ", now.Monday())              // 2013-11-18 00:00:00 Mon
	fmt.Println("Sunday: ", now.Sunday())             // 2013-11-24 00:00:00 Sun (Next Sunday)
	fmt.Println("End of Sunday: ", now.EndOfSunday()) // 2013-11-24 23:59:59.999999999 Sun (End of next Sunday)

	t = time.Date(2013, 11, 24, 17, 51, 49, 123456789, time.Now().Location()) // 2013-11-24 17:51:49.123456789 Sun
	fmt.Println("Change Date====================================================")
	fmt.Println("Monday: ", now.New(t).Monday())             // 2013-11-18 00:00:00 Sun (Last Monday if today is Sunday)
	fmt.Println("Sunday: ", now.New(t).Sunday())             // 2013-11-24 00:00:00 Sun (Beginning Of Today if today is Sunday)
	fmt.Println("End of Sunday: ", now.New(t).EndOfSunday()) // 2013-11-24 23:59:59.999999999 Sun (End of Today if today is Sunday)
}
