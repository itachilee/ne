package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Rdb    *redis.Client
	MgoCli *mongo.Client
	CosCli *cos.Client
)
