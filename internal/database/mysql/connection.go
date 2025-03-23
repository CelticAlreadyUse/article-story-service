package mysqldb

import (
	"fmt"

	"github.com/CelticAlreadyUse/article-story-service/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func InitConnectionString()string{
connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	config.DBUSER(),config.DBPass(),config.DBHOST(),config.DBPort(),config.DBName())
	return connStr
}
func MysqlConnection() (*gorm.DB){
	db, err := gorm.Open(mysql.Open(InitConnectionString()), &gorm.Config{})
	if err !=nil{
		panic(err)
	}
	logrus.Info("database connected : "+InitConnectionString())
	return db
}