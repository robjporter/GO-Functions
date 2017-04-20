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
		browser, browserversion, devicename, osname, osversion := getBrowser(c.Request().UserAgent())
		s.Statuses[status]++
		s.Browsers[browser]++
		s.BrowserVersion[browser+" "+browserversion]++
		s.OS[osname]++
		s.OSVersion[osname+" "+osversion]++
		s.Device[devicename]++
		return nil
	}
}

func getBrowser(agent string) (string, string, string, string, string) {
	ua1 := browser.Parse(agent)
	return ua1.Browser.Name, ua1.Browser.Version, ua1.Device.Name, ua1.OS.Name, ua1.OS.Version
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}
