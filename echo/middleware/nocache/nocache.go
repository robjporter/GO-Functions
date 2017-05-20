package nocache

import (
	"time"

	"github.com/labstack/echo"
)

// ServerHeader middleware adds a `Server` header to the response.
func NoCacheHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Expires", time.Unix(0, 0).Format(time.RFC1123))
		c.Response().Header().Set("Cache-Control", "no-cache, private, max-age=0")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("X-Accel-Expires", "0")
		return next(c)
	}
}
