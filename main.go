package main

import (
	"log"
	"net/http"
	"time"

	"github.com/itachilee/gin/global"
	"github.com/itachilee/gin/internal/routers"
	"github.com/itachilee/gin/models"
	"github.com/itachilee/gin/pkg/cache"
	"github.com/itachilee/gin/pkg/setting"
)

func init() {

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
		panic(err)
	}
}

// @title title ne on go
// @version 1.0
// @description mijiu.group description
// @termsOfService http://swagger.io/terms/

// @contact.name leonlee
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, jsonstr)
	// })
	s.ListenAndServe()
}
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Log", &global.LogSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = models.Setup()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = cache.SetupRedis()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Mongo", &global.MongoSetting)
	if err != nil {
		return err
	}

	cache.SetupMongo()
	if err != nil {
		return err
	}
	return nil
}
