package model

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Diary struct {
	Date     string
	Mood     string
	Content  string
	AuthorId int
}

func SaveDiary(Content string, AuthorId int) error {
	diary := new(Diary)
	diary.Content = Content
	diary.AuthorId = AuthorId
	year, month, day := time.Now().Date()
	diary.Date = fmt.Sprintf("%d-%d-%d", year, month, day)
	err := db.Create(&diary).Error
	if err != nil {
		logrus.Error("Failed saving diary to DB,msg:", err.Error())
	}
	return err
}
