package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitLoadWithViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
func DBHOST() string {
	return viper.GetString("mysql.dbhost")
}
func PORT_HTTP() string {
	return viper.GetString("port")
}
func DBUSER() string {
	return viper.GetString("mysql.dbuser")
}
func DBName() string {
	return viper.GetString("mysql.dbname")
}

func DBPass() string {
	return viper.GetString("mysql.dbpass")
}

func DBPort() string {
	return viper.GetString("mysql.dbport")
}
