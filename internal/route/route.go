package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/milezoom/prayer-times/internal/calculation"
	"github.com/milezoom/prayer-times/internal/usecase"
)

// SetRoute put route to echo instan
func SetRoute(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.GET("/dhuhr", func(c echo.Context) error {
		doy := calculation.GetDayOfYear()
		ad := calculation.GetAngularDistanceToDate(float64(doy))
		eqTime := calculation.GetEquationOfTime(float64(doy), ad)
		return c.String(http.StatusOK, usecase.CalculateDhuhr(eqTime))
	})

	return e
}
