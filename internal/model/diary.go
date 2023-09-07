package model

import (
	"ChattyDiaryBot/internal/util"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Diary struct {
	Date      string
	Content   string
	TimeStamp int64
	AuthorId  int
}

func SaveDiary(Content string, AuthorId int) error {
	diary := new(Diary)
	diary.Content = util.Encode(Content)
	diary.AuthorId = AuthorId
	year, month, day := time.Now().Date()
	diary.Date = fmt.Sprintf("%d-%d-%d", year, month, day)
	diary.TimeStamp = time.Now().Local().Unix()
	err := db.Create(&diary).Error
	if err != nil {
		logrus.Error("Failed saving diary to DB,msg:", err.Error())
	}
	return err
}

//
//Gei diary after the given UNIX time
//
func SelectDiary(uid int, timestampAfter int64) ([]Diary, error) {
	var diaries []Diary
	err := db.Where("time_stamp > ? AND author_id = ?", timestampAfter, uid).Find(&diaries).Error
	if err != nil {
		logrus.Error("failed querying diary,msg:" + err.Error())
	}
	return diaries, err
}
