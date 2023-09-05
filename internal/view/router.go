package view

import (
	"ChattyDiaryBot/internal/controller"

	"github.com/labstack/echo/v4"
)

func InitReverseListener(e *echo.Echo) {
	e.POST("/Listener", controller.ReverseHttpHandle)
}
