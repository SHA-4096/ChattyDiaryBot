package util

import (
	"ChattyDiaryBot/internal/config"
	param "ChattyDiaryBot/internal/util/params"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type EndpointStruct struct {
	SendMessageEndpoint string
}

var EndPoint EndpointStruct

func init() {
	EndPoint.SendMessageEndpoint = "/send_msg"
}

//
//send content to user or group of id
//messageType:priiveta or group
//
func SendMessage(id int, content, messageType string) error {
	//filling data into struct
	bodyTmp := new(param.SendMessageParam)
	bodyTmp.Auto_escape = false
	if messageType == "private" {
		bodyTmp.User_id = id
	} else if messageType == "group" {
		bodyTmp.Group_id = id
	} else {
		return fmt.Errorf("please specify a valid messageType")
	}
	bodyTmp.Message = content
	bodyTmp.Message_type = messageType

	//posting
	cli := new(http.Client)
	reqUrl := fmt.Sprintf("http://%s:%s%s", config.Config.Server.Addr, config.Config.Server.Port, EndPoint.SendMessageEndpoint)

	resp, err := cli.PostForm(reqUrl,
		url.Values{
			"message":      {bodyTmp.Message},
			"message_type": {bodyTmp.Message_type},
			"group_id":     {fmt.Sprint(bodyTmp.Group_id)},
			"user_id":      {fmt.Sprint(bodyTmp.User_id)},
			"auto_escape":  {fmt.Sprint(bodyTmp.Auto_escape)},
		})

	if err != nil {
		logrus.Error("Error when posting,msg:", err.Error())
	}
	body := make([]byte, 512)
	n, err := resp.Body.Read(body)
	logrus.Info("Got response from server,msg=", string(body[:n]))
	return err
}
