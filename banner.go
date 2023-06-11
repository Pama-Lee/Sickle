package main

import (
	"github.com/fatih/color"
)

var (
	banner = `
__                      _                _
\ \    /\  /\___   ___ | | _____ _ __   | |
 \ \  / /_/ / _ \ / _ \| |/ / _ \ '__| / __)
 / / / __  / (_) | (_) |   <  __/ |    \__ \
/_/  \/ /_/ \___/ \___/|_|\_\___|_|    (   /
                                        |_|
`
)

func printBanner() {
	// 背景浅蓝色, 字体荧光白, 输出完毕后恢复默认
	c := color.New(color.BgBlue, color.FgHiWhite).Add(color.Bold)
	_, err := c.Println(banner)
	if err != nil {
		return
	}
	c = color.New(color.Reset)
	_, _ = c.Println("")
}
