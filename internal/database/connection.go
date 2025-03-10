package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnection () (*gorm.DB){
	dsn := "user:#Wahyu123@tcp(127.0.0.1:3306)/storis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		panic(err)
	}
	logrus.Info("database connected %s",dsn)
	return db
}