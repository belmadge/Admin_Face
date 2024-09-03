package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	CPF       string    `gorm:"size:14"`
	Passport  string    `gorm:"size:50"`
	Role      string    `gorm:"size:50"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
