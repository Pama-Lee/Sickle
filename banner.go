package main

import (
	"github.com/fatih/color"
)

var (
	banner = `
__       _____ _      __   __         __
\ \     / ___/(_)____/ /__/ /__     _/ /
 \ \    \__ \/ / ___/ //_/ / _ \   / __/
 / /   ___/ / / /__/ ,< / /  __/  (_  )
/_/   /____/_/\___/_/|_/_/\___/  /  _/
                                 /_/
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
