package model

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Color     string
}
