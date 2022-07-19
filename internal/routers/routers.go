package routers

import (
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	ginMiddlewares "github.com/Laisky/gin-middlewares"
	docs "github.com/itachilee/gin/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/itachilee/gin/global"
	"github.com/itachilee/gin/graph"
	"github.com/itachilee/gin/graph/generated"
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
	// r.Use(middleware.LoggerToMongo())
	r.Use(gin.Recovery())
	// r.Use(middleware.BasicAuth())
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

	// r.GET("/", http.PlaygroundHandler())
	// r.POST("/query", http.GraphQLHandler())
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		MgoCli: global.MgoCli,
		CosCli: global.CosCli,
	}}))
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.MultipartForm{})
	r.GET("/", ginMiddlewares.FromStd(playground.Handler("GraphQL playground", "/query")))
	r.POST("/query", ginMiddlewares.FromStd(h.ServeHTTP))
	// r.GET("/query", func(c *gin.Context){
	// 	srv
	// })
	return r
}
