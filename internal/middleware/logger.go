package middleware

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itachilee/gin/global"
	"github.com/itachilee/gin/internal/hook"
	"github.com/sirupsen/logrus"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := global.LogSetting.LogFilePath
	logFileName := global.LogSetting.LogFileName

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err := os.Stat(fileName)

	if err != nil {
		fmt.Println(err)

	}

	f, err := os.Create(fileName)

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	if err != nil {
		fmt.Println(err.Error())
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	// logger.SetFormatter(&logrus.TextFormatter{})

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

var once sync.Once

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {

	// //实例化
	logger := logrus.New()
	once.Do(func() {

		hooker, err := hook.NewHookerWithAuth(
			global.MongoSetting.MgoUrl,
			global.MongoSetting.Db,
			global.MongoSetting.Collection,
			global.MongoSetting.User,
			global.MongoSetting.Pass)
		if err == nil {
			logger.Hooks.Add(hooker)
		} else {
			fmt.Print(err)
		}
	})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_referer"] = c.Request.Referer()
		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		accessLogMap["request_client_ip"] = c.ClientIP()

		accessLogMap["response_time"] = endTime
		// 日志格式
		logger.WithFields(accessLogMap).Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
