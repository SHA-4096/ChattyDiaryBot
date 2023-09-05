package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	Host      HostStruct      `yaml:"host"`
	LogConfig logConfigStruct `yaml:"logConfig"`
}

type HostStruct struct {
	Port string `yaml:"port"`
}

type logConfigStruct struct {
	LogPath    string `yaml:"logPath"`
	RotateTime int    `yaml:"rotateTime"`
	MaxAge     int    `yaml:"maxAge"`
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
