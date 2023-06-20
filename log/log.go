package log

import (
	"Sickle/entity"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
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
	prefixColor := color.New(color.BgHiBlue).Add(color.Bold).Add(color.FgHiWhite).SprintFunc()
	log.SetPrefix(prefixColor("["+config.Project.Name+"]") + " ")
	// 设置日志输出位置
	if config.Log.File != "" {
		// 设置日志文件, 文件名: `config.Log.File.时间戳.log` 格式
		// 保存在Log目录下
		// 检查目录是否存在, 不存在则创建
		Dir := "logs"
		if !FileExists(Dir) {
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

// GinLogger 自定义日志记录中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 在请求完成后记录日志
		if c.Writer.Status() == 404 {
			codeString := color.New(color.BgHiYellow).Add(color.FgHiWhite).SprintFunc()
			Warn(c.Request.Method, c.Request.URL, codeString(c.Writer.Status()), c.Writer.Size(), c.ClientIP())
		}
		if c.Writer.Status() == 500 {
			codeString := color.New(color.BgHiRed).Add(color.FgHiWhite).SprintFunc()
			Error(c.Request.Method, c.Request.URL, codeString(c.Writer.Status()), c.Writer.Size(), c.ClientIP())
		}
		if c.Writer.Status() == 200 {
			// 计算耗时时间
			stopTime := time.Now()
			esTime := stopTime.Sub(startTime)
			elapsedString := color.New(color.FgHiYellow).SprintFunc()
			codeString := color.New(color.BgHiGreen).Add(color.FgHiWhite).SprintFunc()
			Info(codeString(c.Writer.Status()), c.Request.Method, c.Request.URL, elapsedString(esTime), c.ClientIP())
		}
	}
}

// 处理args参数
func handleArgs(args ...interface{}) string {
	// 在不同param中间加空格
	var result []string
	for _, arg := range args {
		result = append(result, fmt.Sprintf("%v", arg))
	}

	param := strings.Join(result, " ")
	return param
}

// Debug 打印调试日志
func Debug(args ...interface{}) {
	if ServerConfig.Log.Level != "debug" {
		return
	}
	debugString := color.New(color.FgHiBlue).SprintFunc()
	c := fmt.Sprintf("%s %s", debugString("[DEBUG]"), handleArgs(args...))
	log.Println(c)
}

// Info 打印信息日志
func Info(args ...interface{}) {
	if ServerConfig.Log.Level != "debug" && ServerConfig.Log.Level != "info" {
		return
	}
	infoString := color.New(color.FgHiGreen).SprintFunc()

	c := fmt.Sprintf("%s %s", infoString("[INFO]"), handleArgs(args...))
	log.Println(c)
}

// Warn 打印警告日志
func Warn(args ...interface{}) {
	if ServerConfig.Log.Level != "debug" && ServerConfig.Log.Level != "info" && ServerConfig.Log.Level != "warn" {
		return
	}
	warnString := color.New(color.FgHiYellow).SprintFunc()
	c := fmt.Sprintf("%s %s", warnString("[WARN]"), handleArgs(args...))
	log.Println(c)
}

// Error 打印错误日志
func Error(args ...interface{}) {
	errorString := color.New(color.FgHiRed).SprintFunc()
	c := fmt.Sprintf("%s %s", errorString("[ERROR]"), handleArgs(args...))
	log.Println(c)
}

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	// 打开文件
	_, err := os.Open(path)
	// 文件不存在
	if err != nil && os.IsNotExist(err) {
		return false
	}
	// 文件存在
	return true
}
