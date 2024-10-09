package main

import (
	"github.com/iamchristiancollins/weather-backend/config"
	"github.com/iamchristiancollins/weather-backend/db"
	"github.com/iamchristiancollins/weather-backend/models"
	"github.com/iamchristiancollins/weather-backend/routes"
)

func main() {
	config.LoadConfig()
	db.Connect()
	db.DB.AutoMigrate(&models.User{}, &models.Rating{})
	router := routes.SetupRouter()
	router.Run(":" + config.AppConfig.Server.Port)
}
