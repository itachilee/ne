package global

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Rdb    *redis.Client
	MgoCli *mongo.Client
)
