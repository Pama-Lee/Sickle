package main

import (
	"Hooker/database"
	"Hooker/log"
	"os"
)

/**
   ██░ ██  ▒█████   ▒█████   ██ ▄█▀▓█████  ██▀███
  ▓██░ ██▒▒██▒  ██▒▒██▒  ██▒ ██▄█▒ ▓█   ▀ ▓██ ▒ ██▒
  ▒██▀▀██░▒██░  ██▒▒██░  ██▒▓███▄░ ▒███   ▓██ ░▄█ ▒
  ░▓█ ░██ ▒██   ██░▒██   ██░▓██ █▄ ▒▓█  ▄ ▒██▀▀█▄
  ░▓█▒░██▓░ ████▓▒░░ ████▓▒░▒██▒ █▄░▒████▒░██▓ ▒██▒
   ▒ ░░▒░▒░ ▒░▒░▒░ ░ ▒░▒░▒░ ▒ ▒▒ ▓▒░░ ▒░ ░░ ▒▓ ░▒▓░
   ▒ ░▒░ ░  ░ ▒ ▒░   ░ ▒ ▒░ ░ ░▒ ▒░ ░ ░  ░  ░▒ ░ ▒░
   ░  ░░ ░░ ░ ░ ▒  ░ ░ ░ ▒  ░ ░░ ░    ░     ░░   ░
   ░  ░  ░    ░ ░      ░ ░  ░  ░      ░  ░   ░


__                      _                _
\ \    /\  /\___   ___ | | _____ _ __   | |
 \ \  / /_/ / _ \ / _ \| |/ / _ \ '__| / __)
 / / / __  / (_) | (_) |   <  __/ |    \__ \
/_/  \/ /_/ \___/ \___/|_|\_\___|_|    (   /
                                        |_|

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
	log.Info("Hooker is exiting...")
	// 退出程序
	os.Exit(code)
}
