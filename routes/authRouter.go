package routes

import (
	"go-jwt-authentication/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/users/refresh-token", controllers.RefreshToken())
	incomingRoutes.POST("/users/logout", controllers.Logout())
}
