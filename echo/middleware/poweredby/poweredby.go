package poweredby

import "github.com/labstack/echo"

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Powered-By", "Echo/3.1.0")
		return next(c)
	}
}
