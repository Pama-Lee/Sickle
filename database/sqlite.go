package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitSqlite 初始化sqlite数据库
func InitSqlite() error {
	// 使用gorm连接数据库
	db, err := gorm.Open(sqlite.Open("hooker.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	// 将数据库连接赋值给全局变量
	Database = db

	return nil
}
