package view

import (
	"ChattyDiaryBot/internal/controller"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {
	e.POST("/Listener", controller.ReverseHttpHandle)
}
