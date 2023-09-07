package controller

import (
	param "ChattyDiaryBot/internal/controller/params"
	"ChattyDiaryBot/internal/model"
	"ChattyDiaryBot/internal/util"
	"fmt"
	"strings"

	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ReverseHttpHandle(e echo.Context) error {
	var msg param.MessagePrivate
	err := e.Bind(&msg)
	if err != nil {
		logrus.Error("Error handling reverse http,msg:" + err.Error())
	}
	if msg.Message_id == 0 {
		return nil
	}
	//update state if needed
	res, _ := regexp.MatchString("/d", msg.Message)
	if res {
		param.SetUserState(fmt.Sprintf("%d", msg.User_id), "1")
	}
	//

	state, _ := param.GetUserState(fmt.Sprintf("%d", msg.User_id))

	if state == "1" {
		content := strings.Trim(msg.Message, "/d")
		err := model.SaveDiary(content, msg.User_id)
		if err != nil {
			util.SendMessage(msg.User_id, "啊哦，bot好像似了……", "private")
			return err
		}
		util.SendMessage(msg.User_id, "好的，今天的日记已保存！", "private")
		param.SetUserState(fmt.Sprintf("%d", msg.User_id), "0")
	}
	defer e.Request().Body.Close()
	fmt.Printf("%d:%s", msg.User_id, msg.Message)
	logrus.Info("Received message:" + msg.Raw_message)
	return nil
}
