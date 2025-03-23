package mongodb

import (
	"context"
	"fmt"

	"github.com/CelticAlreadyUse/article-story-service/internal/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnstringMongo() string {
	config.InitLoadWithViper()
	return fmt.Sprintf("mongodb://%s:%s", config.MongoHOST(), config.MongoPort(),)
}
func Connect() *mongo.Database {
	config.InitLoadWithViper()
	clientOption := options.Client().ApplyURI(ConnstringMongo())
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		logrus.Warn("Mongodb failed to connect err:", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logrus.Warn(err)
	} else {
		logrus.Info("Sucessfully connected to mongoDB")
	} 
	return client.Database(config.MongoDBName())
}
