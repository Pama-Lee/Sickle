package database

// Forward 转发Webhooks的实体
type Forward struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Url       string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Enable    bool   `gorm:"type:tinyint(1);not null;default:1"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

// ForwardLog 转发Webhooks的日志实体
type ForwardLog struct {
	ID        uint   `gorm:"primarykey"`
	ForwardID uint   `gorm:"type:int(11);not null;index"`
	Body      string `gorm:"type:longtext;not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

// ForwardLogDetail 转发Webhooks的日志详情实体
type ForwardLogDetail struct {
	ID           uint   `gorm:"primarykey"`
	ForwardLogID uint   `gorm:"type:int(11);not null;index"`
	Succeed      bool   `gorm:"type:tinyint(1);not null;default:1"`
	RetryCount   uint   `gorm:"type:int(11);not null;default:0"`
	ResponseBody string `gorm:"type:longtext;not null"`
	UpdatedAt    int64  `gorm:"autoUpdateTime"`
	CreatedAt    int64  `gorm:"autoCreateTime"`
}
