package database

import (
	"Sickle/log"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type VersionedEntityRoot struct {
	Entities []VersionedEntity
	Name     string
	Desc     string
}

type VersionedEntity struct {
	Entity  interface{}
	Version string
}

// SetUp 映射数据库实体
func SetUp() {
	db := Database

	entities := []VersionedEntityRoot{
		{
			Name:     "Bot",
			Desc:     "Bot实体",
			Entities: GetBotEntity(),
		},
		{
			Name:     "Forward",
			Desc:     "Forward实体",
			Entities: GetForwardEntity(),
		},
		{
			Name:     "Webhook",
			Desc:     "Webhook实体",
			Entities: GetWebhookEntity(),
		},
	}

	// 获取entity包下的所有实体
	for _, entity := range entities {
		log.Debug(fmt.Sprintf("Migrate %s", entity.Name))
		for _, versionedEntity := range entity.Entities {
			err := Migrate(db, versionedEntity.Entity, versionedEntity.Version)
			if err != nil {
				log.Error(fmt.Sprintf("Migrate %s - [%s] failed: %s", entity.Name, versionedEntity.Version, err.Error()))
			}
		}
	}

}

// Migrate 迁移数据库
func Migrate(db *gorm.DB, entity interface{}, version string) error {
	err := db.AutoMigrate(entity)
	if err != nil {
		return err
	}

	entityType := reflect.TypeOf(entity).Elem()
	log.Debug(fmt.Sprintf("Migrate %s - [%s] successfully", entityType.Name(), version))

	return nil
}
