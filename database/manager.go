package database

import (
	"Sickle/entity"
	"Sickle/log"
	"errors"
	"gorm.io/gorm"
)

var (
	// Database 数据库
	Database *gorm.DB
)

// InitDatabase 初始化数据库
func InitDatabase(config *entity.Config) error {
	// 检查数据库类型
	err := checkDatabaseConfig(config)
	if err != nil {
		return err
	}
	log.Debug("Check databaseEntity config successfully")
	// 初始化数据库
	if config.Database.Type == "mysql" {
		err := InitMysql(config)
		if err != nil {
			return err
		}
	}

	if config.Database.Type == "sqlite" {
		err := InitSqlite()
		if err != nil {
			return err
		}
	}

	log.Info("Use databaseEntity: " + config.Database.Type + " successfully")

	// 映射数据库实体
	SetUp()

	return nil
}

// checkDatabaseConfig 检查数据库配置
func checkDatabaseConfig(config *entity.Config) error {
	log.Debug("Check databaseEntity config")
	DatabaseConfig := config.Database
	// 检查其中的字段是否为空
	if DatabaseConfig.Type == "" {
		return errors.New("the databaseEntity type is empty")
	}
	switch DatabaseConfig.Type {
	case "mysql":
		if DatabaseConfig.Type == "mysql" {
			// 检查数据库地址是否为空
			if DatabaseConfig.Host == "" {
				return errors.New("the databaseEntity host is empty")
			}
			// 检查数据库端口是否为空
			if DatabaseConfig.Port == 0 {
				return errors.New("the databaseEntity port is empty")
			}
			// 检查数据库用户名是否为空
			if DatabaseConfig.User == "" {
				return errors.New("the databaseEntity user is empty")
			}
			// 检查数据库密码是否为空
			if DatabaseConfig.Password == "" {
				return errors.New("the databaseEntity password is empty")
			}
			// 检查数据库名称是否为空
			if DatabaseConfig.Database == "" {
				return errors.New("the databaseEntity name is empty")
			}
			return nil
		}
		break
	case "sqlite":
		if DatabaseConfig.Type == "sqlite" {
			return nil
		}
		break
	default:
		return errors.New("the databaseEntity type is not supported")
	}
	return errors.New("unknown error")
}

// GetDatabaseSession 获取数据库会话
func GetDatabaseSession() *gorm.DB {
	return Database.Session(&gorm.Session{})
}

// CloseDatabase 关闭数据库
func CloseDatabase() error {
	// 关闭数据库
	db, err := Database.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}
