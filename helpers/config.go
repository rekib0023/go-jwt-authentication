package helpers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SECRET_KEY  string
}

var AppConfig Config

func LoadConfig(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	config := Config{
		PORT:        port,
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     db_port,
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		SECRET_KEY:  os.Getenv("SECRET_KEY"),
	}

	AppConfig = config
}
