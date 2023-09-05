package main

import (
	"ChattyDiaryBot/internal/config"
	"ChattyDiaryBot/internal/util"
	"ChattyDiaryBot/internal/view"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	config.InitConfig()
	e := echo.New()
	e.GET("/state", func(c echo.Context) error {
		return c.String(http.StatusOK, "The server is running")
	})
	logrus.Info("Setting up logrus")
	util.InitLogrus()
	logrus.Info("Setting up Routers")
	view.InitReverseListener(e)
	logrus.Info("Starting network service")
	e.Logger.Fatal(e.Start(":" + config.Config.Host.Port))
}
