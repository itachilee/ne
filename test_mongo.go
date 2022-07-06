package main

import (
	"context"
	"fmt"

	"github.com/itachilee/gin/models"
	"github.com/itachilee/gin/pkg/cache"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func test_mongo() {

	var (
		client     = cache.GetMgoCli()
		err        error
		collection *mongo.Collection

		iResult *mongo.InsertOneResult
		id      primitive.ObjectID
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("blog").Collection("log")
	lr := &models.User{Id: 2, Name: "test1"}
	//插入某一条数据
	if iResult, err = collection.InsertOne(context.TODO(), lr); err != nil {
		fmt.Print(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())

	return
}
