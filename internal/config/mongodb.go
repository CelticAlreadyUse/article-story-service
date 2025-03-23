package config

import (

	"github.com/spf13/viper"
)

func MongoHOST() string{
	return viper.GetString("mongodb.dbhost")
}
func MongoUSER() string{
	return viper.GetString("mongodb.dbuser")
}
func MongoDBName() string{
	return viper.GetString("mongodb.dbname")
}
func MongoPort() string{
	return viper.GetString("mongodb.dbport")
}
func MongoCollection() string{
	return viper.GetString("mongodb.dbcollection")
}