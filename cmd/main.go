package main

import (
	"ChattyDiaryBot/internal/config"
	"ChattyDiaryBot/internal/model"
	"ChattyDiaryBot/internal/util"
	"ChattyDiaryBot/internal/view"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("Loading Config")
	config.InitConfig()

	logrus.Info("Setting up logrus")
	util.InitLogrus()

	logrus.Info("Connecting to Database")
	model.InitRedisDB()
	model.InitMysql()

	e := echo.New()
	e.GET("/state", func(c echo.Context) error {
		return c.String(http.StatusOK, "The server is running")
	})

	logrus.Info("Setting up Routers")
	view.SetRouters(e)
	logrus.Info("Starting network service")
	e.Logger.Fatal(e.Start(":" + config.Config.Host.Port))
}
