package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamchristiancollins/weather-backend/models"
	"github.com/iamchristiancollins/weather-backend/services"
	"net/http"
)

type RatingController {
	ratingService *services.RatingService
}

func SubmitRating(c *gin.Context) {
	var input models.RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetString("userID")

    rating, err := services.CreateRating(userID, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"rating": rating})
}
}