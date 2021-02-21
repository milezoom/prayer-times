package main

import (
	"github.com/labstack/echo/v4"
	"github.com/milezoom/prayer-times/internal/config"
	"github.com/milezoom/prayer-times/internal/route"
)

func main() {
	e := echo.New()

	route.SetRoute(e)
	config.SetConfig(e)

	e.Logger.Fatal(e.Start(":" + config.GetPort()))
}
