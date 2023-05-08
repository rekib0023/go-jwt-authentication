package controllers

import (
	"go-jwt-authentication/database"
	"go-jwt-authentication/models"
	"go-jwt-authentication/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func UserResponse(user models.User) serializers.UserSerializer {
	return serializers.UserSerializer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existing_user models.User
		var user models.User

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		database.Database.Db.Find(&existing_user, "email = ?", user.Email)

		if existing_user.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists with this email address"})
			return
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

		user.Password = password

		database.Database.Db.Create(&user)

		responseUser := UserResponse(user)

		c.JSON(http.StatusCreated, responseUser)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]string

		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if data["email"] == "" || data["password"] == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide email and password"})
			return
		}

		var user models.User

		database.Database.Db.Where("email = ?", data["email"]).First(&user)

		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No user found with this email"})
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
			return
		}
	}
}
