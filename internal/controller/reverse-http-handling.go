package controller

import (
	param "ChattyDiaryBot/internal/controller/params"
	"ChattyDiaryBot/internal/model"
	"ChattyDiaryBot/internal/util"
	cq "ChattyDiaryBot/internal/util/cq-code"
	"fmt"
	"strconv"
	"time"

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
	//check if a goroutine is running for this user
	//if so,return
	state, err := param.GetUserState(fmt.Sprintf("%d", msg.User_id))
	if err != nil {
		logrus.Error("Error getting user state,msg:" + err.Error())
	}
	if state == "-1" {
		go util.SendMessage(msg.User_id, "请等待上一次操作完成喵~", "private")
		return nil
	}
	//update state if needed
	res, _ := regexp.MatchString("/r", msg.Message)
	if res {
		param.SetUserState(fmt.Sprintf("%d", msg.User_id), "2")
	}
	res, _ = regexp.MatchString("/d", msg.Message)
	if res {
		param.SetUserState(fmt.Sprintf("%d", msg.User_id), "1")
	}
	res, _ = regexp.MatchString("/s", msg.Message)
	if res {
		param.SetUserState(fmt.Sprintf("%d", msg.User_id), "3")
	}

	state, _ = param.GetUserState(fmt.Sprintf("%d", msg.User_id))

	if state == "1" {
		go storeDiary(msg)
	} else if state == "2" {
		go queryDiary(msg)
	} else if state == "3" {
		go getImage(msg)
	}
	defer e.Request().Body.Close()
	fmt.Printf("%d:%s", msg.User_id, msg.Message)
	logrus.Info("Received message:" + msg.Raw_message)
	return nil
}

//=============Below are subfunc of the controller functions

func storeDiary(msg param.MessagePrivate) error {
	param.SetUserState(fmt.Sprintf("%d", msg.User_id), "-1")
	content := msg.Message[3:]
	err := model.SaveDiary(content, msg.User_id)
	if err != nil {
		util.SendMessage(msg.User_id, "啊哦，bot好像似了……", "private")
		return err
	}
	util.SendMessage(msg.User_id, "好的，今天的日记已保存喵！", "private")
	param.SetUserState(fmt.Sprintf("%d", msg.User_id), "0")
	return nil
}

func queryDiary(msg param.MessagePrivate) error {
	param.SetUserState(fmt.Sprintf("%d", msg.User_id), "-1")
	timeBack, err := strconv.Atoi(msg.Message[3:])
	if err != nil || (timeBack <= 0 || (time.Now().Unix()-int64(timeBack*3600*24) < 0)) {
		util.SendMessage(msg.User_id, "我认不出这个时间喵！", "private")
		return fmt.Errorf("not a valid day count input")
	}
	diaries, _ := model.SelectDiary(msg.User_id, time.Now().Unix()-int64(timeBack)*3600*24)
	content := fmt.Sprintf("这是你近%d天的日记喵~\n\n", timeBack)
	for _, diary := range diaries {
		content += fmt.Sprintf("[%s]\n%s\n\n", diary.Date, util.Decode(diary.Content))
	}
	util.SendMessage(msg.User_id, content, "private")
	logrus.Info(fmt.Sprintf("Finished pushing diary to %d", msg.User_id))
	param.SetUserState(fmt.Sprintf("%d", msg.User_id), "0")
	return nil

}

func getImage(msg param.MessagePrivate) error {
	param.SetUserState(fmt.Sprintf("%d", msg.User_id), "-1")
	content := msg.Message[3:]
	url, err := util.BingImageSearch(content)
	if err != nil {
		logrus.Error("Bing image API not successful,msg:" + err.Error())
	}
	util.SendMessage(msg.User_id, cq.MarshalImage(url), "private")
	param.SetUserState(fmt.Sprintf("%d", msg.User_id), "0")
	return nil

}
