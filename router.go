package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "work-wechat/config"
	"work-wechat/middleware"
	"work-wechat/service"
	message_queue "work-wechat/service/message-queue"
	"work-wechat/service/practice"
	"work-wechat/service/rpc"
)

/**
 * @Description: 路由初始化
 * @return *gin.Engine
 */
func initRouter() *gin.Engine {

	// debug
	gin.SetMode(App.Logger.Debug)

	// 命令行颜色
	//if App.Logger.ConsoleColor {
	//	gin.DisableConsoleColor()
	//}

	router := gin.New()

	rpcRouter := router.Group("rpc")
	{
		// 启动 rpc server
		rpcRouter.GET("server", func(context *gin.Context) {
			rpc.RunRpcServer()
		})

		// 调用方法
		rpcRouter.GET("call", func(context *gin.Context) {
			//rpc.CallRpcFunc()
			rpc.GoFunc()
		})
	}

	// 测试
	router.GET("/test", func(context *gin.Context) {
		practice.Exec()
	})

	router.GET("/enqueue", func(context *gin.Context) {
		message_queue.Enqueue()
	})

	router.GET("/consumer", func(context *gin.Context) {
		message_queue.Consumer()
	})

	router.Use(middleware.LoggerToFile(), gin.Recovery())
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "welcome")
	})

	workRouter := router.Group("work", func(c *gin.Context) {
		c.Writer.Header().Add("content-type", "application/json; charset=utf-8")
	})
	// 打卡记录
	workRouter.POST("punch", service.WxApp.GetPunchRecord)
	// 打卡日报
	workRouter.POST("/punch/daily", service.WxApp.GetDailyPunch)
	// 打开月报
	workRouter.POST("/punch/month", service.WxApp.GetMonthPunch)
	return router
}
