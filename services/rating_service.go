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

func GetRatingsByUserID(userID uint) ([]models.Rating, error) {
	var ratings []models.Rating
	query := db.DB.Where("user_id = ?", userID)

	// Execute the query and store the result
	result := query.Find(&ratings)
	// Check for errors
	if result.Error != nil {
		return nil, result.Error
	}
	// Return the ratings and no error
	return ratings, nil
}
