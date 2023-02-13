package log

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLog() {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
		return
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	//实例化
	Logger = logrus.New()

	//设置输出
	Logger.Out = src

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

}
