package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-gin-api/config"
	"go-gin-api/util"
	"os"
	"path"
	"time"
)

var Logger = logrus.New()


func InitLogBash(logFilePath string, logLevel logrus.Level, writeMap lfshook.WriterMap) *logrus.Logger {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	if ok, _ := util.PathExists(logFilePath); !ok {
		// Directory not exist
		fmt.Println("Create " + logFilePath)
		_ = os.Mkdir(logFilePath, os.ModePerm)
	}
	//设置输出
	Logger.Out = src

	// 设置日志级别
	Logger.SetLevel(logLevel)

	lfhook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05",})

	//新增hook
	Logger.AddHook(lfhook)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{})

	return Logger
}

func setRotatelog(fileName string) *rotatelogs.RotateLogs  {
	logWriter, _ := rotatelogs.New(
		//分割文件后的文件名
		fileName+"-%Y-%m-%d.log",
		//生成如软链接，指向最新的日志文件
		rotatelogs.WithLinkName(fileName),
		//设置最大的保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志的切割时间间隔
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return logWriter
}

func InitLog() *logrus.Logger  {
	fmt.Println("init service ....")
	fileName := path.Join(config.GinConfig.Log.Path, config.GinConfig.Log.Name)
	logWriter := setRotatelog(fileName)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel:  logWriter,
		logrus.FatalLevel:  logWriter,
		logrus.PanicLevel:  logWriter,
	}
	return InitLogBash(config.GinConfig.Log.Path, logrus.InfoLevel, writeMap)
}

func InitWarnLog() *logrus.Logger  {
	fileName := path.Join(config.GinConfig.Log.ErrorPath, config.GinConfig.Log.ErrorName)
	logWriter := setRotatelog(fileName)
	writeMap := lfshook.WriterMap{
		logrus.WarnLevel:  logWriter,
	}
	return InitLogBash(config.GinConfig.Log.ErrorPath, logrus.WarnLevel, writeMap)
}