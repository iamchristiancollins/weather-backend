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