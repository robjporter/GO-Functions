package times

import (
	"time"

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
	return dat + " " + tim // + " " + t.timezone.String()
}

func (t *Times) formattedDate() (time.Time, error) {
	//ti, _ := time.Parse(t.format, t.buildTimeString())
	//fmt.Println(ti.Format("02-01-2006 15:04:05 -0700"))
	//return ti.Format(t.format), err
	return time.Parse(t.format, t.buildTimeString())
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

///////////////// TO IMPLEMENT /////////////////
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

func (t *Times) IsFuture() bool {
	return false
}

func (t *Times) IsPast() bool {
	return false
}

func (t *Times) Difference() string {
	return "false"
}

func (t *Times) DiffInSeconds() string {
	return "false"
}

func (t *Times) DiffInMinutes() string {
	return "false"
}

func (t *Times) DiffInHours() string {
	return "false"
}

func (t *Times) DiffInDays() string {
	return "false"
}

func (t *Times) DiffInWeeks() string {
	return "false"
}

func (t *Times) DiffInMonths() string {
	return "false"
}

func (t *Times) DiffInYears() string {
	return "false"
}

func (t *Times) InRFC1889() (time.Time, error) {
	return time.Now(), nil
}

///////////////// IMPLEMENTED /////////////////
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
	t.hour = 23
	t.minute = 59
	t.second = 59
	return t.formattedDate()
}

func (t *Times) EndOfWeek() (time.Time, error) {
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, _ := t.formattedDate()
	tmp2 := 7 - t.getWeekDayNumber(tmp.Weekday())
	diff := (t.getWeekDayNumber(tmp.Weekday()) - tmp2) + 1
	return tmp.AddDate(0, 0, diff), nil

}

func (t *Times) EndOfMonth() (time.Time, error) {
	t.day = t.maxDaysInMonth()
	t.hour = 23
	t.minute = 59
	t.second = 59
	return t.formattedDate()
}

func (t *Times) EndOfYear() (time.Time, error) {
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	return t.formattedDate()
}

func (t *Times) EndOfDecade() (time.Time, error) {
	tmp := as.ToString(t.year)
	tmp = tmp[0:3] + "9"
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	t.year = as.ToInt(tmp)
	return t.formattedDate()
}

func (t *Times) EndOfCentury() (time.Time, error) {
	tmp := as.ToString(t.year)
	tmp = tmp[0:2] + "99"
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	t.year = as.ToInt(tmp)
	return t.formattedDate()
}

func (t *Times) StartOfCentury() (time.Time, error) {
	tmp := as.ToString(t.year)
	tmp = tmp[0:2] + "00"
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	t.year = as.ToInt(tmp)
	return t.formattedDate()
}

func (t *Times) StartOfDecade() (time.Time, error) {
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	t.year = t.year - (t.year % 10)
	return t.formattedDate()
}

func (t *Times) StartOfYear() (time.Time, error) {
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	return t.formattedDate()
}

func (t *Times) StartOfDay() (time.Time, error) {
	t.hour = 00
	t.minute = 00
	t.second = 00
	return t.formattedDate()

}

func (t *Times) StartOfWeek() (time.Time, error) {
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()
	tmp2 := t.getWeekDayNumber(tmp.Weekday()) - 1
	diff := t.getWeekDayNumber(tmp.Weekday()) - tmp2
	return tmp.AddDate(0, 0, -diff), nil
}

func (t *Times) StartOfMonth() (time.Time, error) {
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()
	return tmp.AddDate(0, 0, -(t.day)+1), nil
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
		return tmp.AddDate(1, 0, 0), nil
	}
	return tmp, err
}

func (t *Times) AddYears(year int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(year, 0, 0), nil
	}
	return tmp, err
}

func (t *Times) AddMonth() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 1, 0), nil
	}
	return tmp, err
}

func (t *Times) AddMonths(month int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, month, 0), nil
	}
	return tmp, err
}

func (t *Times) AddWeek() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, 7), nil
	}
	return tmp, err
}

func (t *Times) AddWeeks(week int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, week*7), nil
	}
	return tmp, err
}

func (t *Times) AddDay() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, 1), nil
	}
	return tmp, err
}

func (t *Times) AddDays(day int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, day), nil
	}
	return tmp, err
}

func (t *Times) AddHour() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(1 * time.Hour)), nil
	}
	return tmp, err
}

func (t *Times) AddHours(hour int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(time.Duration(hour) * time.Hour)), nil
	}
	return tmp, err
}

func (t *Times) AddMinute() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(1 * time.Minute)), nil
	}
	return tmp, err
}

func (t *Times) AddMinutes(minute int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(time.Duration(minute) * time.Minute)), nil
	}
	return tmp, err
}

func (t *Times) AddSecond() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(1 * time.Second)), nil
	}
	return tmp, err
}

func (t *Times) AddSeconds(second int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(time.Duration(second) * time.Second)), nil
	}
	return tmp, err
}

func (t *Times) SubYear() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(-1, 0, 0), nil
	}
	return tmp, err
}

func (t *Times) SubYears(year int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(-year, 0, 0), nil
	}
	return tmp, err
}

func (t *Times) SubMonth() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, -1, 0), nil
	}
	return tmp, err
}

func (t *Times) SubMonths(month int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, -month, 0), nil
	}
	return tmp, err
}

func (t *Times) SubWeek() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, -7), nil
	}
	return tmp, err
}

func (t *Times) SubWeeks(week int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, -(week * 7)), nil
	}
	return tmp, err
}
func (t *Times) SubDay() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, -1), nil
	}
	return tmp, err
}

func (t *Times) SubDays(day int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.AddDate(0, 0, -day), nil
	}
	return tmp, err
}

func (t *Times) SubHour() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(-1 * time.Hour)), nil
	}
	return tmp, err
}

func (t *Times) SubHours(hour int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(time.Duration(-hour) * time.Hour)), nil
	}
	return tmp, err
}

func (t *Times) SubMinute() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(-1 * time.Minute)), nil
	}
	return tmp, err
}

func (t *Times) SubMinutes(minute int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(time.Duration(-minute) * time.Minute)), nil
	}
	return tmp, err
}

func (t *Times) SubSecond() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(-1 * time.Second)), nil
	}
	return tmp, err
}

func (t *Times) SubSeconds(second int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		return tmp.Add(time.Duration(time.Duration(-second) * time.Second)), nil
	}
	return tmp, err
}
