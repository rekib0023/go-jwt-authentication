package main

import (
	"go-jwt-authentication/database"
	"go-jwt-authentication/helpers"
	"go-jwt-authentication/middleware"
	"go-jwt-authentication/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.LoadConfig(".env")

	AppConfig := helpers.AppConfig

	database.Connect()

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Use(middleware.AuthMiddleware())

	routes.UserRoutes(router)

	router.Run(":" + AppConfig.PORT)
}
