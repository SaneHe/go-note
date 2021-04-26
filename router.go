package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "work-wechat/config"
	"work-wechat/middleware"
	"work-wechat/service"
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
