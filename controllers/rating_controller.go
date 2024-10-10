package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamchristiancollins/weather-backend/models"
	"github.com/iamchristiancollins/weather-backend/services"
	"net/http"
)

func SubmitRating(c *gin.Context) {
	var input models.RatingInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error with authentication"})
		return
	}

	rating, err := services.CreateRating(uint(userID.(float64)), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"rating": rating})
}

func GetRatings(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	ratings, err := services.GetRatingsByUserID(uint(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ratings": ratings})
}
