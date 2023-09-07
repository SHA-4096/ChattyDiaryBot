package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

//
//Set a key-value pair
//
func SetKeyValuePair(key, value string) error {
	err := redisClient.Do(redisCtx, "set", key, value).Err()
	if err != nil {
		logrus.Error(fmt.Sprintf("Error at model.SetKeyValuePair,msg=%s", err.Error()))
		return err
	}
	return nil
}

//
//Set key expiration time in redis
//
func SetExpiration(key string, seconds int) error {
	err = redisClient.Do(redisCtx, "expire", key, seconds).Err()
	if err != nil {
		logrus.Error(fmt.Sprintf("Error at model.SetExpiration,msg=%s", err.Error()))
		return err
	}
	return nil
}

//
//Discard key expiration time in redis
//
func DiscardExpiration(key string) error {
	err = redisClient.Do(redisCtx, "persist", key).Err()
	if err != nil {
		logrus.Error(fmt.Sprintf("Error at model.DiscardExpiration,msg=%s", err.Error()))
		return err
	}
	return nil
}

//
//Get the value of a specific key in redis
//
func GetKeyValue(key string) (string, error) {
	val, err := redisClient.Get(redisCtx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil

}
