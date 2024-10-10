package services

import (
	"github.com/iamchristiancollins/weather-backend/db"
	"github.com/iamchristiancollins/weather-backend/models"
)

func CreateRating(userID uint, input models.RatingInput) (models.Rating, error) {
	rating := models.Rating{
		UserID:       userID,
		Temperature:  input.Temperature,
		TopRating:    input.TopRating,
		BottomRating: input.BottomRating,
	}

	if err := db.DB.Create(&rating).Error; err != nil {
		return models.Rating{}, err
	}

	return rating, nil
}
