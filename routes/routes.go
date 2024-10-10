package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamchristiancollins/weather-backend/middlewares"
	"github.com/iamchristiancollins/weather-backend/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/user", controllers.GetUser)
		authorized.POST("/ratings", controllers.SubmitRating)
		authorized.GET("/recommendations", controllers.GetRecommendations)
	}

	return router
}
