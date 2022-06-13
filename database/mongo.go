package database

import (
	"context"
	"fmt"
	"go-line-demo/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb *mongo.Client

func Init() {
	// 取得mongo db連線
	uri := assembleUri()
	mongodb, _ = mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := mongodb.Connect(ctx)

	defer func() {
		if err != nil {
			panic(err)
		}
	}()
}

func assembleUri() string {
	settings := config.GetConfig()
	user := settings.Get("DB_USERNAME")
	pwd := settings.Get("DB_PASS")
	host := settings.Get("DB_HOST")
	port := settings.Get("DB_PORT")

	return fmt.Sprintf("mongodb://%v:%v@%v:%v", user, pwd, host, port)
}

func GetClient() *mongo.Client {
	return mongodb
}

func Close() {
	if err := mongodb.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
