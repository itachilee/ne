package routers

import (
	"time"

	"github.com/gin-gonic/gin"

	docs "github.com/itachilee/gin/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/itachilee/gin/internal/middleware"
	v1 "github.com/itachilee/gin/internal/routers/api/v1"
	"github.com/itachilee/gin/pkg/limiter"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {

	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
	r := gin.New()
	// r.Use(middleware.LoggerToFile())
	r.Use(middleware.LoggerToMongo())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.ContextTimeout(30 * time.Second))
	r.Use(middleware.RateLimiter(methodLimiters))
	user := v1.NewUser()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ding", v1.Helloworld)

		apiv1.POST("/users", user.Create)
		apiv1.DELETE("/users/:id", user.Delete)
		apiv1.PUT("/users/:id", user.Update)
		apiv1.PATCH("/users/:id/status", user.Update)
		apiv1.GET("/users", user.List)
	}
	return r
}
