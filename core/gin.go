package core

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"strconv"
	"time"
)

var Router *gin.Engine

/*
initGin 初始化gin
*/
func initGin() {
	//创建gin engine
	engine := gin.New()
	//注册全局中间件，使用自定义的日志中间件，使用gin默认的recover中间件
	engine.Use(ginLogger(), gin.Recovery())
	//404处理
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, response.Response{StatusCode: 404, StatusMsg: "Not Found!"})
	})
	Router = engine
}

/*
ginLogger 是自定义的全局日志中间件，用于替代gin的默认日志中间件
*/
func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start) // 本次请求的总共消耗时间
		// 写入日志
		global.Logger.Infow(
			strconv.Itoa(c.Writer.Status()),
			"path", c.Request.URL.String(),
			"method", c.Request.Method,
			"clientIp", c.ClientIP(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"cost", cost,
		)
	}
}
