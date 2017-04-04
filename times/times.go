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
	loc, _ := time.LoadLocation("Europe/London")
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
}

func (t *Times) IsLeapYear() bool {
	return t.year%4 == 0 && (t.year%100 != 0 || t.year%400 == 0)
}

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
	return time.Parse(t.format, t.buildTimeString())
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
