package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ReverseHttpHandle(e echo.Context) error {
	var msg MessagePrivate
	err := e.Bind(&msg)
	if err != nil {
		logrus.Error("Error handling reverse http,msg:" + err.Error())
	}
	if msg.Message_id == 0 {
		return nil
	}
	defer e.Request().Body.Close()
	logrus.Info(fmt.Sprintf("Incoming message from %s", e.Request().Host))
	fmt.Printf("%d:%s", msg.User_id, msg.Message)
	logrus.Info("Received message:" + msg.Raw_message)
	return nil
}
