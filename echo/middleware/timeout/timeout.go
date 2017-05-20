package nocache

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var timeout time.Duration

func Timeout(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		a := c.(context.Context)
		ctx, cancel := context.WithTimeout(a, timeout)
		defer func() {
			cancel()
			if ctx.Err() == context.DeadlineExceeded {
				c.Response().WriteHeader(http.StatusGatewayTimeout)
			}
		}()
		return next(c)
	}
}

func New(timeout time.Duration) {
	timeout = timeout
}
