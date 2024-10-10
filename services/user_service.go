package services

import (
	"errors"
	"github.com/iamchristiancollins/weather-backend/db"
	"github.com/iamchristiancollins/weather-backend/models"
	"github.com/iamchristiancollins/weather-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(input models.RegisterInput) (models.User, error) {
	//hash password and save user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func AuthenticateUser(input models.LoginInput) (string, error) {
	//authenticate user credentials
	var user models.User
	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return "", errors.New("incorrect username or password")
	}
	query := db.DB.Where("username = ?", input.Username)
	result := query.Find(&user)

	if result.Error != nil {
		return "", errors.New("incorrect username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("incorrect username or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserByID(userID string) (models.User, error) {
	var user models.User

	query := db.DB.Where("id = ?", userID)
	result := query.First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
