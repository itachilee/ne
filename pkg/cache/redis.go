package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/itachilee/gin/global"
)

var (
	ctx  = context.Background()
	once sync.Once
)

func SetupRedis() error {
	var err error

	opt := &redis.Options{
		Addr:        global.RedisSetting.Addr,
		Password:    global.RedisSetting.Password,
		IdleTimeout: time.Duration(global.RedisSetting.IdleTimeout),
		DB:          global.RedisSetting.DB,
	}
	// global.Rdb = new(redis.Client)
	once.Do(func() {
		global.Rdb = redis.NewClient(opt)
		err = global.Rdb.Ping(ctx).Err()
		if err != nil {
			fmt.Println(err)
		}
	})
	return err
}

func GetRedisCli() *redis.Client {
	if global.Rdb == nil {
		SetupRedis()
	}
	return global.Rdb
}
