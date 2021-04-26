package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
	"work-wechat/logger"
)

/**
  中间价
*/
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()

		//body, err := ioutil.ReadAll(c.Request.Body)
		//if err != nil {
		//	logger.Error("解析请求体出错:",  err)
		//}
		//
		//header, err := json.Marshal(c.Request.Header)
		//if err != nil {
		//	logger.Error("解析请求头出错:", err)
		//}

		//日志格式
		logger.Info(logrus.Fields{
			// 状态码
			"http_status": c.Writer.Status(),
			// 执行时间
			"total_time": fmt.Sprintf("%6v", time.Now().Sub(startTime)),
			// 客户端 ip
			"ip": c.ClientIP(),
			// 请求方式
			"method": c.Request.Method,
			// 请求路径
			"uri": c.Request.RequestURI,
			// 请求头
			//"header":  header,
			//
			//// 请求体
			//"body": body,

			//"response": c.Writer.,
		})
	}
}
