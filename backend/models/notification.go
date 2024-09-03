package models

import (
	"time"
)

type Notification struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Message   string `gorm:"type:text"`
	Type      string `gorm:"size:50"`
	SentAt    time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
