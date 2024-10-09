package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	Temperature  int    `gorm:"not null"`
	TopRating    string `gorm:"not null"`
	BottomRating string `gorm:"not null"`
}
