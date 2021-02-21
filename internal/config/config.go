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
	e.Logger.SetOutput(utils.LogWriter())
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Output:  utils.LogWriter(),
		Format:  `[${time_rfc3339}] ${method} ${uri} ${error}`,
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
