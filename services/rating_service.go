package services

import (
	"github.com/yourusername/yourproject-backend/db"
	"github.com/yourusername/yourproject-backend/models"
)

func CreateRating(userID string, input models.RatingInput) (models.Rating, error) {
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
