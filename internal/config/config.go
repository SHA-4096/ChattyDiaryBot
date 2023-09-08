package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	Host      HostStruct      `yaml:"host"`
	Server    ServerStruct    `yaml:"server"`
	LogConfig logConfigStruct `yaml:"logConfig"`
	Database  DatabaseStruct  `yaml:"database"`
	Bot       BotStruct       `yaml:"bot"`
	Users     []int           `yaml:"users"`
}

type DatabaseStruct struct {
	Mysql MysqlStruct `yaml:"mysql"`
	Redis RedisStruct `yaml:"redis"`
}

type HostStruct struct {
	Port string `yaml:"port"`
}

type ServerStruct struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

type logConfigStruct struct {
	LogPath    string `yaml:"logPath"`
	RotateTime int    `yaml:"rotateTime"`
	MaxAge     int    `yaml:"maxAge"`
}

type MysqlStruct struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbName"`
	Timeout  string `yaml:"timeout"`
}

type RedisStruct struct {
	RedisAddr     string `yaml:"redisAddr"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDB       int    `yaml:"redisDB"`
}

type BotStruct struct {
	RemindHour int    `yaml:"remindHour"`
	BingAPIKey string `yaml:"bingAPIKey"`
}

var Config ConfigStruct

func InitConfig() {
	conf, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		logrus.Panic("Failed to Read config,err=" + err.Error())
		panic("Failed to read config")
	}
	yaml.Unmarshal(conf, &Config)
	logrus.Info("Config Loaded")

}
