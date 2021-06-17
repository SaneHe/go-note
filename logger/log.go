package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"work-wechat/config"
)

var logger = logrus.New()

/**
 * @Description: 初始化 log
 */
func init() {
	level, err := logrus.ParseLevel(config.App.Logger.Debug)
	if err != nil {
		Fatal("debug 转换失败，请检查")
	}

	logger.SetLevel(level)
	logFileWriter, err := os.OpenFile(config.App.Logger.LogFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Fatal("创建 static 文件失败", err)
	}
	defer logger.Writer().Close()

	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: config.App.Logger.LogTimestampFormat,
	}
	logger.SetReportCaller(config.App.Logger.ReportCaller)
	logger.SetOutput(io.MultiWriter(logFileWriter, os.Stdout))
}

/**
 * @Description: 日志记录 debug
 * @param args
 */
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

/**
 * @Description: 日志记录 info
 * @param args
 */
func Info(args ...interface{}) {
	logger.Info(args...)
}

/**
 * @Description: 日志记录 error
 * @param args
 */
func Error(args ...interface{}) {
	logger.Error(args...)
}

/**
 * @Description: 日志记录 trace
 * @param args
 */
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

/**
 * @Description: 日志 fatal
 * @param args
 */
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
