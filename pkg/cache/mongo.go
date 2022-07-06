package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/itachilee/gin/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo() {
	var err error
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		global.MongoSetting.User,
		global.MongoSetting.Pass,
		global.MongoSetting.MgoUrl,
		global.MongoSetting.Db,
	)
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到MongoDB
	global.MgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = global.MgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func GetMgoCli() *mongo.Client {
	if global.MgoCli == nil {
		SetupMongo()
	}
	return global.MgoCli
}
