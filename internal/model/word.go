package model

import (
	"time"
)

type Word struct {
	ID         uint       `gorm:"primaryKey"`
	Word       string     `gorm:"not null"`
	Categories []Category `gorm:"many2many:word_categories;"`
	UserID     string     `gorm:"not null"`
	SavedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Lists      []List     `gorm:"many2many:word_lists;"`
}
