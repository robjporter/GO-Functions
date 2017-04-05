package times

import (
	"time"

	"strings"

	"github.com/robjporter/go-functions/as"
)

type Times struct {
	year     int
	month    int
	day      int
	hour     int
	minute   int
	second   int
	timezone *time.Location
	err      error
	format   string
	reset    bool
}

type Diff struct {
	year  int
	month int
	week  int
	day   int
	hour  int
	min   int
	sec   int
}

func New(year, month, day, hour, minute, second int, location string) *Times {
	loc, err := time.LoadLocation(location)
	if err == nil {

		t := Times{
			year:     year,
			month:    month,
			day:      day,
			hour:     hour,
			minute:   minute,
			second:   second,
			timezone: loc,
			err:      nil,
			format:   "2006-01-02 15:04:05",
			reset:    true,
		}
		return &t
	} else {
		return nil
	}
}

///////////////// HELPERS /////////////////
func processNumber(num string) string {
	if len(num) == 1 {
		return "0" + num
	}
	return num
}

func (t *Times) buildTimeString() string {
	dat := as.ToString(t.year) + "-" + processNumber(as.ToString(t.month)) + "-" + processNumber(as.ToString(t.day))
	tim := as.ToString(t.hour) + ":" + processNumber(as.ToString(t.minute)) + ":" + processNumber(as.ToString(t.second))
	return dat + " " + tim
}

func (t *Times) formattedDate() (time.Time, error) {
	return time.Parse(t.format, t.buildTimeString())
}

func MonthNameToNumber(month time.Month) int {
	switch month {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return 0
}

func (t *Times) getWeekDayNumber(day time.Weekday) int {
	switch day {
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	case time.Sunday:
		return 7
	}
	return 0
}

func (t *Times) maxDaysInMonth() int {
	switch t.month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if t.IsLeapYear() {
			return 29
		} else {
			return 28
		}
	}
	return 0
}

func (t *Times) updateValues(date time.Time) {
	t.year = date.Year()
	t.month = MonthNameToNumber(date.Month())
	t.day = date.Day()
	t.hour = date.Hour()
	t.minute = date.Minute()
	t.second = date.Second()
}

func (t *Times) formatDifference(data time.Duration) Diff {
	tmp := Diff{}
	splits := strings.Split(as.ToString(data), "h")
	tmp.hour = 0
	if len(splits) > 1 {
		tmp.hour = as.ToInt(splits[0])
	}
	splits = strings.Split(as.ToString(data), "m")
	splits2 := strings.Split(splits[0], "h")
	tmp.min = 0
	if len(splits2) > 1 {
		tmp.min = as.ToInt(splits2[1])
	}
	tmp.sec = 0
	splits = strings.Split(as.ToString(data), "s")
	splits2 = strings.Split(splits[0], "m")
	if len(splits2) > 1 {
		tmp.sec = as.ToInt(splits2[1])
	}

	tmp.day = 0
	if tmp.hour > 23 {
		tmp.day = tmp.hour / 24
		tmp.hour = tmp.hour % 24
	}

	tmp.week = 0
	if tmp.day > 6 {
		tmp.week = tmp.day / 7
		tmp.day = tmp.day % 7
	}

	tmp.month = 0
	if tmp.week > 4 {
		tmp.month = tmp.week / 4
		tmp.week = tmp.week % 4
	}

	tmp.year = 0
	if tmp.month > 12 {
		tmp.year = tmp.month / 12
		tmp.month = tmp.month % 12
	}

	return tmp
}

func (t *Times) processStruct(tmp Diff) string {
	result := ""
	if tmp.sec > 0 {
		result = as.ToString(tmp.sec) + "s"
	}
	if tmp.min > 0 {
		result = as.ToString(tmp.min) + "m" + result
	}
	if tmp.hour > 0 {
		result = as.ToString(tmp.hour) + "h" + result
	}
	if tmp.day > 0 {
		result = as.ToString(tmp.day) + "d" + result
	}
	if tmp.week > 0 {
		result = as.ToString(tmp.week) + "w" + result
	}
	if tmp.month > 0 {
		result = as.ToString(tmp.month) + "M" + result
	}
	if tmp.year > 0 {
		result = as.ToString(tmp.year) + "y" + result
	}
	return result
}

///////////////// TO IMPLEMENT /////////////////

///////////////// IMPLEMENTED /////////////////
func (t *Times) Difference(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) DifferenceDiff(t2 *Times) Diff {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) DiffInSeconds(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Seconds())
}

func (t *Times) DiffInMinutes(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Minutes())
}

func (t *Times) DiffInHours(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Hours())
}

func (t *Times) DiffInDays(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Hours() / 24)
}

func (t *Times) DiffInWeeks(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString((tmp3.Hours() / 24) / 7)
}

func (t *Times) DiffInMonths(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(((tmp3.Hours() / 24) / 7) / 4)
}

func (t *Times) DiffInYears(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString((((tmp3.Hours() / 24) / 7) / 4) / 12)
}

func (t *Times) IsFuture(t2 *Times) bool {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	if tmp2.After(tmp) {
		return true
	}

	return false
}

func (t *Times) IsPast(t2 *Times) bool {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	if tmp.Before(tmp2) {
		return true
	}

	return false
}

func (t *Times) ISOWeek() (string, error) {
	tmp, _ := t.formattedDate()
	_, week := tmp.ISOWeek()
	return as.ToString(week), nil
}

func (t *Times) Format1() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Jan")
	return result, err
}

func (t *Times) Format2() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05 PM Jan")
	return result, err
}

func (t *Times) Format3() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/Jan/2006 03:04:05 PM")
	return result, err
}

func (t *Times) Format4() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/Jan/2006 15:04:05")
	return result, err
}

func (t *Times) Format5() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Mon Jan")
	return result, err
}

func (t *Times) Format6() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Monay January")
	return result, err
}

func (t *Times) Format7() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Jan")
	return result, err
}

func (t *Times) Format8() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("2/1/6 3:4:5 PM")
	return result, err
}

func (t *Times) Format9() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("_2/1/6 3:4:5 PM")
	return result, err
}

func (t *Times) Format10() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM")
	return result, err
}

func (t *Times) Format11() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05 PM")
	return result, err
}

func (t *Times) Format12() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05.000 PM")
	return result, err
}

func (t *Times) Format13() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05.000000 PM")
	return result, err
}

func (t *Times) Format14() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05.000000000 PM")
	return result, err
}

func (t *Times) Format15() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 MST")
	return result, err
}

func (t *Times) Format16() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z7")
	return result, err
}

func (t *Times) Format17() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z07")
	return result, err
}

func (t *Times) Format18() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z0700")
	return result, err
}

func (t *Times) Format19() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z07:00")
	return result, err
}

func (t *Times) Format20() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 -07:00")
	return result, err
}

func (t *Times) Format822() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC822)
	return result, err
}

func (t *Times) Format1123() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC1123)
	return result, err
}

func (t *Times) Format1123z() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC1123Z)
	return result, err
}

func (t *Times) Format3339() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC3339)
	return result, err
}

func (t *Times) Format3339nano() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC3339Nano)
	return result, err
}

func (t *Times) Format8222z() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC822Z)
	return result, err
}

func (t *Times) Format850() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC850)
	return result, err
}

func (t *Times) TimePrevious(day time.Weekday) (time.Time, error) {
	tmp, _ := t.formattedDate()
	currentDay := t.getWeekDayNumber(tmp.Weekday())
	pastDay := t.getWeekDayNumber(day)
	days := 0

	if currentDay > pastDay {
		days = currentDay - pastDay
	} else if currentDay == pastDay {
		days = 7
	} else {
		days = 7 + (currentDay - pastDay)
	}

	return tmp.AddDate(0, 0, -days), nil
}

func (t *Times) TimeNext(day time.Weekday) (time.Time, error) {
	tmp, _ := t.formattedDate()
	currentDay := t.getWeekDayNumber(tmp.Weekday())
	futureDay := t.getWeekDayNumber(day)
	days := 0

	if currentDay < futureDay {
		days = futureDay - currentDay
	} else if currentDay == futureDay {
		days = 7
	} else {
		days = 7 + (futureDay - currentDay)
	}

	return tmp.AddDate(0, 0, days), nil
}

func (t *Times) EndOfDay() (time.Time, error) {
	t2 := *t
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfWorkWeek() (time.Time, error) {
	t2 := *t
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()

	diff := 5 - t.getWeekDayNumber(tmp.Weekday())

	tmp = tmp.AddDate(0, 0, diff)
	if t.reset {
		*t = t2
	}
	return tmp, err

}

func (t *Times) EndOfWeek() (time.Time, error) {
	t2 := *t
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()

	diff := 7 - t.getWeekDayNumber(tmp.Weekday())

	tmp = tmp.AddDate(0, 0, diff)
	if t.reset {
		*t = t2
	}
	return tmp, err

}

func (t *Times) EndOfMonth() (time.Time, error) {
	t2 := *t
	t.day = t.maxDaysInMonth()
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfYear() (time.Time, error) {
	t2 := *t
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfDecade() (time.Time, error) {
	t2 := *t
	tmp2 := as.ToString(t.year)
	tmp2 = tmp2[0:3] + "9"
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	t.year = as.ToInt(tmp2)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfCentury() (time.Time, error) {
	t2 := *t
	tmp2 := as.ToString(t.year)
	tmp2 = tmp2[0:2] + "99"
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	t.year = as.ToInt(tmp2)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfCentury() (time.Time, error) {
	t2 := *t
	tmp2 := as.ToString(t.year)
	tmp2 = tmp2[0:2] + "00"
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	t.year = as.ToInt(tmp2)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfDecade() (time.Time, error) {
	t2 := *t
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	t.year = t.year - (t.year % 10)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfYear() (time.Time, error) {
	t2 := *t
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfDay() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfWorkWeek() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()

	diff := t.getWeekDayNumber(tmp.Weekday()) - 1

	tmp = tmp.AddDate(0, 0, -diff)
	if t.reset {
		*t = t2
	}
	return tmp, nil
}

func (t *Times) StartOfWeek() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()

	diff := t.getWeekDayNumber(tmp.Weekday()) - 1

	tmp = tmp.AddDate(0, 0, -diff)
	if t.reset {
		*t = t2
	}
	return tmp, nil
}

func (t *Times) StartOfMonth() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()
	tmp = tmp.AddDate(0, 0, -(t.day)+1)
	if t.reset {
		*t = t2
	}
	return tmp, nil
}

func (t *Times) IsWeekend() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Saturday || ti.Weekday() == time.Sunday {
		return true
	}
	return false
}

func (t *Times) IsWorkday() bool {
	if t.IsWeekend() {
		return false
	}
	return true
}

func (t *Times) IsMonday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Monday {
		return true
	}
	return false
}

func (t *Times) IsTuesday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Tuesday {
		return true
	}
	return false
}

func (t *Times) IsWednesday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Wednesday {
		return true
	}
	return false
}

func (t *Times) IsThursday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Thursday {
		return true
	}
	return false
}

func (t *Times) IsFriday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Friday {
		return true
	}
	return false
}

func (t *Times) IsSaturday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Saturday {
		return true
	}
	return false
}

func (t *Times) IsSunday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Sunday {
		return true
	}
	return false
}

func (t *Times) IsLeapYear() bool {
	return t.year%4 == 0 && (t.year%100 != 0 || t.year%400 == 0)
}

func (t *Times) AddYear() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(1, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddYears(year int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(year, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMonth() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 1, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMonths(month int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, month, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddWeek() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, 7)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddWeeks(week int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, week*7)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddDay() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, 1)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddDays(day int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, day)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddHour() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(1 * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddHours(hour int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(hour) * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMinute() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(1 * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMinutes(minute int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(minute) * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddSecond() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(1 * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddSeconds(second int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(second) * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubYear() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(-1, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubYears(year int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(-year, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMonth() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, -1, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMonths(month int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, -month, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubWeek() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -7)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubWeeks(week int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -(week * 7))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubDay() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -1)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubDays(day int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -day)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubHour() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(-1 * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubHours(hour int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(-hour) * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMinute() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(-1 * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMinutes(minute int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(-minute) * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubSecond() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(-1 * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubSeconds(second int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(-second) * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}
