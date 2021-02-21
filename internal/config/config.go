package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" //autoload env
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/milezoom/prayer-times/internal/utils"
)

//SetConfig configure echo instance according to env file
func SetConfig(e *echo.Echo) {
	e.HideBanner = true
	e.Use(middleware.Recover())
	configLog(e)
}

func configLog(e *echo.Echo) {
	e.Logger.SetOutput(utils.LogWriter("debug"))
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${level}")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Output:  utils.LogWriter("access"),
		Format:  `${time_rfc3339} ${method} ${status} ${uri} ${error}` + "\n",
	}))
}

//GetPort return port specified in env file
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
