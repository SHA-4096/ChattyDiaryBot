package middleware

import (
	"ChattyDiaryBot/internal/config"
	"ChattyDiaryBot/internal/model"
	"ChattyDiaryBot/internal/util"
	cq "ChattyDiaryBot/internal/util/cq-code"
	"fmt"
	"time"
)

var RoutinePictureKeyword string

func InitCronTasks() {
	RoutinePictureKeyword = "cat"
	go routineSender()
}

func routineSender() {
	for {
		if time.Now().Local().Hour() == config.Config.Bot.RemindHour {
			_, err := model.GetKeyValue("routine.sender.tag")
			if err != nil {
				model.SetKeyValuePair("routine.sender.tag", "1")
				model.SetExpiration("routine.sender.tag", 7200)
				fmt.Println(config.Config.Users[0])
				for _, user := range config.Config.Users {
					content := fmt.Sprintf("主人好呀，今天是%d月%d日，你今天有什么值得记录的东西吗，可以告诉我喵~\n", time.Now().Local().Month(), time.Now().Local().Day())
					url, _ := util.BingImageSearch(RoutinePictureKeyword)
					content += cq.MarshalImage(url)
					util.SendMessage(user, content, "private")
				}

			}
		}
		time.Sleep(time.Minute * time.Duration(5))
	}

}
