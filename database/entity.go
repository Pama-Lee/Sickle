package database

import "Sickle/entity/databaseEntity"

// GetBotEntity 获取bot实体
func GetBotEntity() []VersionedEntity {
	return []VersionedEntity{
		{
			Entity:  &databaseEntity.Bot{},
			Version: "v0.0.1",
		},
		{
			Entity:  &databaseEntity.BotLog{},
			Version: "v0.0.1",
		},
		{
			Entity:  &databaseEntity.WecomBot{},
			Version: "v0.0.1",
		},
	}
}

// GetForwardEntity 获取forward实体
func GetForwardEntity() []VersionedEntity {
	return []VersionedEntity{
		{
			Entity:  &databaseEntity.Forward{},
			Version: "v0.0.1",
		},
		{
			Entity:  &databaseEntity.ForwardLog{},
			Version: "v0.0.1",
		},
		{
			Entity:  &databaseEntity.ForwardLogDetail{},
			Version: "v0.0.1",
		},
	}
}

// GetWebhookEntity 获取webhook实体
func GetWebhookEntity() []VersionedEntity {
	return []VersionedEntity{
		{
			Entity:  &databaseEntity.Webhooks{},
			Version: "v0.0.1",
		},
		{
			Entity:  &databaseEntity.WebhooksLog{},
			Version: "v0.0.1",
		},
		{
			Entity:  &databaseEntity.WebhooksLogDetail{},
			Version: "v0.0.1",
		},
	}
}
