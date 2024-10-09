package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/iamchristiancollins/weather-backend/config"
	"time"
)

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
