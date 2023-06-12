package server

import (
	"Sickle/database"
	"Sickle/entity/databaseEntity"
)

// VerifyUUID 验证UUID
func VerifyUUID(uuid string) bool {
	db := database.Database
	var webhooks databaseEntity.Webhooks
	db.Where("uuid = ?", uuid).First(&webhooks)
	if webhooks.ID == 0 {
		return false
	}
	return true
}
