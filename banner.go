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
	// 背景浅蓝色, 字体荧光白
	c := color.New(color.BgBlue, color.FgHiWhite)
	_, err := c.Println(banner)
	if err != nil {
		return
	}
}
