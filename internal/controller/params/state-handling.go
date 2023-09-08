package controller

import (
	"ChattyDiaryBot/internal/model"
	"fmt"

	"github.com/sirupsen/logrus"
)

//
//Get the state of a user(Identified by uid)from redisDB,used for chat content control
//none or 0:user not in operation(no pair stored in redisDB)
//1:user require to make a daily log
//
func GetUserState(uid string) (string, error) {
	res, err := model.GetKeyValue(fmt.Sprintf("%s.state", uid))
	if err != nil {
		logrus.Info("If redis does not found record,should be normal" + err.Error())
		return "0", err
	}
	return res, nil

}

//
//Set the state of a user(Identified by uid)in redisDB,used for chat content control
//none:user not in operation(no pair stored in redisDB)
//1:user require to make a daily log
//2:provide user with their latest n days of diary
//3:get an image from bing with the word specified
//
func SetUserState(uid, state string) error {
	err := model.SetKeyValuePair(fmt.Sprintf("%s.state", uid), state)
	if err != nil {
		return err
	}
	err = model.SetExpiration(fmt.Sprintf("%s.state", uid), 3600) //Set expiration in order to save resources
	return err
}
