package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamchristiancollins/weather-backend/services"
	"net/http"
)

func GetWeatherData(c *gin.Context) {
	location := c.Query("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location parameter is required"})
		return
	}

	weatherData, err := services.GetWeatherData(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"weather": weatherData})
}
