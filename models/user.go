package models

import (
	"gorm.io/gorm"
)

// User represents a registered user in the application.
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	// Additional fields can be added here.
}

// RegisterInput represents the expected input for user registration.
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginInput represents the expected input for user login.
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
