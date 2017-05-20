package delay

import (
	"time"

	"github.com/labstack/echo"
	"github.com/robjporter/go-functions/as"
)

var timeout string

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Add-Delay", as.ToString(timeout))

		if timeout != "" {
			delayDuration, err := time.ParseDuration(timeout)

			if err == nil {
				time.Sleep(delayDuration)
			}
		}

		return next(c)
	}
}

func New(timeout string) {
	timeout = timeout
}
