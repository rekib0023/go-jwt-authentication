package main

import (
	"go-jwt-authentication/database"
	"go-jwt-authentication/helpers"
	"go-jwt-authentication/middleware"
	"go-jwt-authentication/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	helpers.LoadConfig(".env")

	AppConfig := helpers.AppConfig

	database.Connect()

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Use(middleware.AuthMiddleware())

	routes.UserRoutes(router)

	router.Run(":" + AppConfig.PORT)
}
