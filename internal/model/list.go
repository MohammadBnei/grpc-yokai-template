package model

type List struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	Words  []Word `gorm:"many2many:word_words"`
	UserID string `gorm:"not null"`
}
