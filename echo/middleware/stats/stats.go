package stats

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/robjporter/go-functions/browser"
)

/*

{
    "total_response_time": "1.907382ms",
    "average_response_time": "86.699\u00b5s",
    "average_response_time_sec": 8.6699e-05,
    "count": 1,
    "pid": 99894,
    "status_code_count": {
        "200": 1
    },
    "time": "2015-03-06 17:23:27.000677896 +0100 CET",
    "total_count": 22,
    "total_response_time_sec": 0.0019073820000000002,
    "total_status_code_count": {
        "200": 22
    },
    "unixtime": 1425659007,
    "uptime": "4m14.502271612s",
    "uptime_sec": 254.502271612
}

*/

type (
	Stats struct {
		Uptime         time.Time      `json:"uptime"`
		RequestCount   uint64         `json:"requestCount"`
		Statuses       map[string]int `json:"statuses"`
		Browsers       map[string]int `json:"browsers"`
		BrowserVersion map[string]int `json:"browserversion"`
		OS             map[string]int `json:"os"`
		Device         map[string]int `json:"device"`
		DeviceType     map[string]int `json:"devicetype"`
		OSVersion      map[string]int `json:"osversion"`
		mutex          sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:         time.Now(),
		Statuses:       map[string]int{},
		Browsers:       map[string]int{},
		BrowserVersion: map[string]int{},
		OS:             map[string]int{},
		Device:         map[string]int{},
		DeviceType:     map[string]int{},
		OSVersion:      map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		browser, browserversion, devicename, devicetype, osname, osversion := getBrowser(c.Request().UserAgent())
		s.Statuses[status]++
		s.Browsers[browser]++
		s.BrowserVersion[browser+" "+browserversion]++
		s.OS[osname]++
		s.OSVersion[osname+" "+osversion]++
		s.Device[devicename]++
		s.DeviceType[devicetype]++
		return nil
	}
}

func getBrowser(agent string) (string, string, string, string, string, string) {
	name := ""
	version := ""
	dname := ""
	osname := ""
	osversion := ""
	devicetype := ""
	ua1 := browser.Parse(agent)

	if ua1.Browser != nil {
		if ua1.Browser.Name != "" {
			name = ua1.Browser.Name
		}
		if ua1.Browser.Version != "" {
			version = ua1.Browser.Version
		}
	} else {
		name = "Unknown"
		version = "Unknown"
	}
	if ua1.Device != nil {
		if ua1.Device.Name != "" {
			dname = ua1.Device.Name
		}
	} else {
		dname = "Unknown"
	}
	if ua1.OS != nil {
		if ua1.OS.Name != "" {
			osname = ua1.OS.Name
		}
		if ua1.OS.Version != "" {
			osversion = ua1.OS.Version
		}
	} else {
		osname = "Unknown"
		osversion = "Unknown"
	}
	if ua1.DeviceType != nil {
		devicetype = ua1.DeviceType.Name
	} else {
		devicetype = "Unknown"
	}
	return name, version, dname, devicetype, osname, osversion
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}
