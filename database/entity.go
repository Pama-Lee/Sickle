package database

import "Sickle/entity/database"

// GetBotEntity 获取bot实体
func GetBotEntity() []VersionedEntity {
	return []VersionedEntity{
		{
			Entity:  &database.Bot{},
			Version: "v0.0.1",
		},
		{
			Entity:  &database.BotLog{},
			Version: "v0.0.1",
		},
		{
			Entity:  &database.WecomBot{},
			Version: "v0.0.1",
		},
	}
}

// GetForwardEntity 获取forward实体
func GetForwardEntity() []VersionedEntity {
	return []VersionedEntity{
		{
			Entity:  &database.Forward{},
			Version: "v0.0.1",
		},
		{
			Entity:  &database.ForwardLog{},
			Version: "v0.0.1",
		},
		{
			Entity:  &database.ForwardLogDetail{},
			Version: "v0.0.1",
		},
	}
}

// GetWebhookEntity 获取webhook实体
func GetWebhookEntity() []VersionedEntity {
	return []VersionedEntity{
		{
			Entity:  &database.Webhooks{},
			Version: "v0.0.1",
		},
		{
			Entity:  &database.WebhooksLog{},
			Version: "v0.0.1",
		},
		{
			Entity:  &database.WebhooksLogDetail{},
			Version: "v0.0.1",
		},
	}
}
