package main

import (
	"Sickle/database"
	"Sickle/log"
	"Sickle/server"
	"os"
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

	// 启动http服务
	server.Start(Config)

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

// Exit 退出程序
func Exit(code int) {
	log.Info("Sickle is exiting...")
	// 退出程序
	os.Exit(code)
}
