package database

import (
	"Sickle/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

// InitMysql 初始化mysql数据库
func InitMysql(config *entity.Config) error {
	// 使用gorm连接数据库
	username := config.Database.User
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port
	portStr := strconv.Itoa(port)
	database := config.Database.Database

	dsn := username + ":" + password + "@tcp(" + host + ":" + portStr + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 将数据库连接赋值给全局变量
	Database = db

	return nil
}
