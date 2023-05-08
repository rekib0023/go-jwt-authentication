package routes

import (
	"go-jwt-authentication/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:userId", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/signin", controllers.Login())

}
