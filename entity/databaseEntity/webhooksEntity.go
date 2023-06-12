package databaseEntity

// Webhooks webhooks实体
type Webhooks struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Url       string `gorm:"type:varchar(255);not null;uniqueIndex"`
	UUID      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Enable    bool   `gorm:"type:tinyint(1);not null;default:1"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

// WebhooksLog webhooks日志实体
type WebhooksLog struct {
	ID         uint   `gorm:"primarykey"`
	WebhooksID uint   `gorm:"type:int(11);not null;index"`
	Body       string `gorm:"type:longtext;not null"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
}

// WebhooksLogDetail webhooks日志详情实体
type WebhooksLogDetail struct {
	ID            uint  `gorm:"primarykey"`
	WebhooksLogID uint  `gorm:"type:int(11);not null;index"`
	UpdatedAt     int64 `gorm:"autoUpdateTime"`
	CreatedAt     int64 `gorm:"autoCreateTime"`
}
