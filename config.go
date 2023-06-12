package main

import (
	"Sickle/entity"
	log2 "Sickle/log"
	"Sickle/tool"
	"encoding/json"
	"log"
	"os"
)

var (
	Config *entity.Config
)

func init() {
	// 初始化配置文件, 程序根目录的`config.json`文件
	// 检查配置文件是否存在, 不存在则创建
	if !tool.FileExists("config.json") {
		// 创建配置文件
		NewConfig()
	}
	// 加载配置文件
	config := LoadConfig()
	// 初始化日志
	log2.InitLogger(config)
	Config = config
}

// LoadConfig 加载配置文件
func LoadConfig() *entity.Config {
	// 打开配置文件
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log2.Error(err)
		}
	}(f)
	// 创建配置文件结构体
	config := &entity.Config{}
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

	config := &entity.Config{}

	// 默认配置
	config.Log.File = "sickle"
	config.Log.Level = "debug"
	config.Server.Host = "127.0.0.1"
	config.Server.Port = 8080
	config.Database.Type = "mysql"
	config.Database.Host = "127.0.0.1"
	config.Database.Port = 3306
	config.Project.Name = "Sickle"
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
