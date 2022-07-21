package routers

import (
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"

	ginMiddlewares "github.com/Laisky/gin-middlewares"
	docs "github.com/itachilee/gin/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/itachilee/gin/global"
	"github.com/itachilee/gin/graph"
	"github.com/itachilee/gin/graph/generated"
	"github.com/itachilee/gin/internal/controller"
	"github.com/itachilee/gin/internal/middleware"
	v1 "github.com/itachilee/gin/internal/routers/api/v1"
	"github.com/itachilee/gin/pkg/limiter"
	"github.com/itachilee/gin/repository"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {

	a, err := gormadapter.NewAdapterByDB(global.Db)
	if err != nil {
		fmt.Print(err)
	}
	e, err := casbin.NewEnforcer("./conf/model.conf", a) // "./conf/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	//add policy
	if hasPolicy := e.HasPolicy("14", "report", "read"); !hasPolicy {
		e.AddPolicy("14", "report", "read")
	}
	if hasPolicy := e.HasPolicy("doctor", "report", "write"); !hasPolicy {
		e.AddPolicy("doctor", "report", "write")
	}
	if hasPolicy := e.HasPolicy("patient", "report", "read"); !hasPolicy {
		e.AddPolicy("patient", "report", "read")
	}
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

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRepository := repository.NewUserRepository(global.Db)
	userController := controller.NewUserController(userRepository)
	apiRoutes := r.Group("/api/v1")

	{
		apiRoutes.GET("/ding", v1.Helloworld)

		apiRoutes.POST("/register", userController.AddUser(e))
		apiRoutes.POST("/signin", userController.SignInUser)
	}
	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", middleware.Authorize("report", "read", e), userController.GetAllUser)
		userProtectedRoutes.GET("/:user", middleware.Authorize("report", "read", e), userController.GetUser)
		userProtectedRoutes.PUT("/:user", middleware.Authorize("report", "write", e), userController.UpdateUser)
		userProtectedRoutes.DELETE("/:user", middleware.Authorize("report", "write", e), userController.DeleteUser)
	}

	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		MgoCli: global.MgoCli,
		CosCli: global.CosCli,
		Db:     global.Db,
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
	return r
}
