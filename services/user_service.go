package services

import (
	"github.com/iamchristiancollins/weather-backend/db"
    "github.com/iamchristiancollins/weather-backend/models"
    "golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	//hash password and save user
	panic("not implemented yet")
}

func AuthenticateUser(username, password string) (models.User, error) {
	//authenticate user credentials
	panic("not implemented yet")
}