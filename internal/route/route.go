package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SetRoute put route to echo instan
func SetRoute(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	return e
}
