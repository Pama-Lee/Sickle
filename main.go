package main

import (
	"Sickle/database"
	"Sickle/log"
	"Sickle/server"
	"Sickle/store"
	"fmt"
	"os"
	"strconv"
)

/**
     ██████  ██▓ ▄████▄   ██ ▄█▀ ██▓    ▓█████
   ▒██    ▒ ▓██▒▒██▀ ▀█   ██▄█▒ ▓██▒    ▓█   ▀
   ░ ▓██▄   ▒██▒▒▓█    ▄ ▓███▄░ ▒██░    ▒███
     ▒   ██▒░██░▒▓▓▄ ▄██▒▓██ █▄ ▒██░    ▒▓█  ▄
   ▒██████▒▒░██░▒ ▓███▀ ░▒██▒ █▄░██████▒░▒████▒
   ▒ ▒▓▒ ▒ ░░▓  ░ ░▒ ▒  ░▒ ▒▒ ▓▒░ ▒░▓  ░░░ ▒░ ░
   ░ ░▒  ░ ░ ▒ ░  ░  ▒   ░ ░▒ ▒░░ ░ ▒  ░ ░ ░  ░
   ░  ░  ░   ▒ ░░        ░ ░░ ░   ░ ░      ░
         ░   ░  ░ ░      ░  ░       ░  ░   ░  ░
                ░
__       _____ _      __   __         __
\ \     / ___/(_)____/ /__/ /__     _/ /
 \ \    \__ \/ / ___/ //_/ / _ \   / __/
 / /   ___/ / / /__/ ,< / /  __/  (_  )
/_/   /____/_/\___/_/|_/_/\___/  /  _/
                                 /_/
*/

func main() {
	printBanner()

	if Config != nil {
		// 初始化数据库
		err := database.InitDatabase(Config)
		if err != nil {
			log.Error(err)
		}

	} else {
		log.Error("Config is nil")
	}

	Test()

	// log.Info(UUID2Config)

	log.Info("Successfully loaded " + strconv.Itoa(len(store.GlobalConfig)) + " webhook(s)")

	// 开启单独的线程处理http
	go server.Start(Config)

	// 开启单独的线程开启控制台
	console()

}

// console 控制台
func console() {
	log.Info("Sickle is running, type 'stop' to stop it")
	for {
		var input string
		// 获取输入
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Error(err)
		}

		switch input {
		case "stop":
			Stop()
		case "exit":
			Exit(0)
		case "reload":
			Reload()
		default:
			log.Warn("Unknown command")
		}
	}
}

// Stop 停止程序
func Stop() {
	// 关闭数据库
	err := database.CloseDatabase()
	if err != nil {
		log.Error(err)
	}
	// 退出程序
	Exit(0)
}

// Reload 重载程序
func Reload() {
	log.Info("Sickle is reloading...")
	// 关闭数据库
	err := database.CloseDatabase()
	if err != nil {
		log.Error(err)
	}
	// 重连数据库
	err = database.InitDatabase(Config)
	if err != nil {
		log.Error(err)
	}
	Test()
	log.Info("Sickle is reloaded")

}

// Exit 退出程序
func Exit(code int) {
	log.Info("Sickle is exiting...")
	// 退出程序
	os.Exit(code)
}
