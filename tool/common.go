package tool

import "os"

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
