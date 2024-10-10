package models

import (
	"gorm.io/gorm"
)

// Rating represents a user's clothing rating for a specific temperature.
type Rating struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	Temperature  int    `gorm:"not null" json:"temperature"`
	TopRating    string `gorm:"not null" json:"top_rating"`
	BottomRating string `gorm:"not null" json:"bottom_rating"`
	// Additional fields can be added here.
}

// RatingInput represents the expected input when submitting a rating.
type RatingInput struct {
	Temperature  int    `json:"temperature" binding:"required"`
	TopRating    string `json:"top_rating" binding:"required"`
	BottomRating string `json:"bottom_rating" binding:"required"`
}
