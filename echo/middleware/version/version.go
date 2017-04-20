package version

import (
	"github.com/labstack/echo"
)

var myVersion = ""

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-App-Version", myVersion)
		return next(c)
	}
}

func New(version string) {
	myVersion = version
}
