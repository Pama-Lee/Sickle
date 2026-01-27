package main

import (
	"Sickle/channel"
	"Sickle/database"
	"Sickle/log"
	"Sickle/server"
	"Sickle/store"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"golang.org/x/term"
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
	hookConfig()
	log.Info("Successfully loaded " + strconv.Itoa(len(store.GlobalConfig)) + " webhook(s)")

	// 开启单独的线程处理channel
	channel.Start()

	// 开启单独的线程处理http
	go server.Start(Config)

	// 检查是否在终端中运行（有 tty）
	if term.IsTerminal(int(os.Stdin.Fd())) {
		// 有终端，开启控制台
		console()
	} else {
		// 没有终端（Docker/后台运行），等待信号
		log.Info("Running in non-interactive mode, use signals to control")
		waitForSignal()
	}

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

	// 如果数据库已经关闭，直接重连
	if database.Database == nil {
		// 重连数据库
		err := database.InitDatabase(Config)
		if err != nil {
			log.Error(err)
		}
	} else {
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
	}

	hookConfig()
	log.Info("Sickle is reloaded")

}

// Exit 退出程序
func Exit(code int) {
	log.Info("Sickle is exiting...")
	// 退出程序
	os.Exit(code)
}

// waitForSignal 等待系统信号（用于 Docker/后台运行模式）
func waitForSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Info("Received shutdown signal")
			Stop()
		case syscall.SIGHUP:
			log.Info("Received reload signal")
			Reload()
		}
	}
}
