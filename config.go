package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config 配置文件结构体
type Config struct {
	// 数据库配置
	Database struct {
		// 数据库类型
		Type string `json:"type"`
		// 数据库地址
		Host string `json:"host"`
		// 数据库端口
		Port int `json:"port"`
		// 数据库用户名
		User string `json:"user"`
		// 数据库密码
		Password string `json:"password"`
		// 数据库名称
		Database string `json:"database"`
	} `json:"database"`
	// 服务器配置
	Server struct {
		// 服务器地址
		Host string `json:"host"`
		// 服务器端口
		Port int `json:"port"`
	} `json:"server"`
	// 日志配置
	Log struct {
		// 日志文件路径
		File string `json:"file"`
		// 日志级别
		Level string `json:"level"`
	} `json:"log"`
	// 项目配置
	Project struct {
		// 项目名称
		Name string `json:"name"`
		// 项目版本
		Version string `json:"version"`
		// 项目作者
		Author string `json:"author"`
		// 项目描述
		Description string `json:"description"`
	} `json:"project"`
}

func init() {
	// 初始化配置文件, 程序根目录的`config.json`文件
	// 检查配置文件是否存在, 不存在则创建
	if !FileExists("config.json") {
		// 创建配置文件
		NewConfig()
	}
	// 加载配置文件
	config := LoadConfig()
	// 初始化日志
	initLogger(config)

}

// LoadConfig 加载配置文件
func LoadConfig() *Config {
	// 打开配置文件
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			Error(err)
		}
	}(f)
	// 创建配置文件结构体
	config := &Config{}
	// 创建JSON解码器
	decoder := json.NewDecoder(f)
	// 解码配置文件
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal(err)
	}
	// 返回配置文件结构体
	return config
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

// NewConfig 创建配置文件
func NewConfig() {
	// 创建配置文件
	f, err := os.Create("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	config := &Config{}

	// 默认配置
	config.Log.File = "hooker"
	config.Log.Level = "debug"
	config.Server.Host = "127.0.0.1"
	config.Server.Port = 8080
	config.Database.Type = "mysql"
	config.Database.Host = "127.0.0.1"
	config.Database.Port = 3306
	config.Project.Name = "Hooker"
	config.Project.Version = "1.0.0"
	config.Project.Author = "Pama Lee"
	config.Project.Description = "A centralized system for managing multi-platform Hooks."

	// 创建JSON编码器
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(config)

	if err != nil {
		log.Fatal(err)
	}
}
