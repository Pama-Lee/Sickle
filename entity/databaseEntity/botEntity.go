package databaseEntity

// Bot bot实体
type Bot struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Url       string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Enable    bool   `gorm:"type:tinyint(1);not null;default:1"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

// WecomBot 企业微信机器人实体
type WecomBot struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Url       string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Enable    bool   `gorm:"type:tinyint(1);not null;default:1"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

// BotLog bot日志实体
type BotLog struct {
	ID        uint   `gorm:"primarykey"`
	BotID     uint   `gorm:"type:int(11);not null;index"`
	BotType   string `gorm:"type:varchar(255);not null;index"`
	Body      string `gorm:"type:longtext;not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}
