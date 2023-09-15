package model

import (
	"ChattyDiaryBot/internal/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
//迁移
//
func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Diary{})
	if err != nil {
		logrus.Panic("Failed to make migrations to mysql.")
		return err
	}
	return nil
}

var db *gorm.DB
var err error

func InitMysql() {
	//dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", config.Config.Database.Mysql.Username, config.Config.Database.Mysql.Password, config.Config.Database.Mysql.Host, config.Config.Database.Mysql.Port, config.Config.Database.Mysql.DbName, config.Config.Database.Mysql.Timeout)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection to mysql failed, error=" + err.Error())
	}
	err = migrate(db)
	if err != nil {
		logrus.Panic("Failed to make migrations to mysql.")
	}
}
