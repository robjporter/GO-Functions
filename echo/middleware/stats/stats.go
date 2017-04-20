package stats

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/robjporter/go-functions/browser"
)

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
