package config

import (
	"fmt"

	"github.com/spf13/viper"
)



func InitLoadWithViper(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") 
	viper.AddConfigPath(".")               
	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func PORT_HTTP() string{
	return viper.GetString("port")
}
func DBName() string{
	return viper.GetString("")
}

func DBPass() string{
	return viper.GetString("")
}

func DBPort() int64{
	return viper.GetInt64("")
}

