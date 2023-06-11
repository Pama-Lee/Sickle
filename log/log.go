package log

import (
	"Hooker/entity"
	"Hooker/tool"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"time"
)

var (
	ServerConfig *entity.Config
)

// Writer 日志输出
var Writer io.Writer

// Timestamp 获取当前时间戳
func Timestamp() int64 {
	// 获取当前时间戳
	return time.Now().Unix()
}

// InitLogger 初始化日志
func InitLogger(config *entity.Config) {
	ServerConfig = config
	// 设置日志前缀
	log.SetPrefix("[" + config.Project.Name + "] ")
	// 设置日志输出位置
	if config.Log.File != "" {
		// 设置日志文件, 文件名: `config.Log.File.时间戳.log` 格式
		// 保存在Log目录下
		// 检查目录是否存在, 不存在则创建
		Dir := "logs"
		if !tool.FileExists(Dir) {
			// 创建目录
			err := os.Mkdir(Dir, 0777)
			if err != nil {
				Error(err)
			}
		}

		LogFileString := fmt.Sprintf("%s.%d.log", config.Log.File, Timestamp())
		logFile, err := os.OpenFile(Dir+"/"+LogFileString, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// 打开日志文件失败, 使用标准输出
			Error(err)
		} else {
			// 使用日志文件, 并且输出到控制台
			log.SetOutput(logFile)
			mw := io.MultiWriter(os.Stdout, logFile)
			log.SetOutput(mw)
			// 设置日志输出
			Writer = mw
		}
	} else {
		Warn("日志文件未配置, 使用标准输出")
	}
}

// Debug 打印调试日志
func Debug(args ...interface{}) {
	if ServerConfig.Log.Level != "debug" {
		return
	}
	debugString := color.New(color.FgHiBlue).SprintFunc()
	c := fmt.Sprintf("%s %s", debugString("[DEBUG]"), fmt.Sprint(args...))
	log.Println(c)
}

// Info 打印信息日志
func Info(args ...interface{}) {
	if ServerConfig.Log.Level != "debug" && ServerConfig.Log.Level != "info" {
		return
	}
	infoString := color.New(color.FgHiGreen).SprintFunc()
	c := fmt.Sprintf("%s %s", infoString("[INFO]"), fmt.Sprint(args...))
	log.Println(c)
}

// Warn 打印警告日志
func Warn(args ...interface{}) {
	if ServerConfig.Log.Level != "debug" && ServerConfig.Log.Level != "info" && ServerConfig.Log.Level != "warn" {
		return
	}
	warnString := color.New(color.FgHiYellow).SprintFunc()
	c := fmt.Sprintf("%s %s", warnString("[WARN]"), fmt.Sprint(args...))
	log.Println(c)
}

// Error 打印错误日志
func Error(args ...interface{}) {
	errorString := color.New(color.FgHiRed).SprintFunc()
	c := fmt.Sprintf("%s %s", errorString("[ERROR]"), fmt.Sprint(args...))
	log.Println(c)
}
