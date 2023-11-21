package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
	DbName     string
}

func LoadConfig(app *AppConfig) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, be ensure all environment variables setting up properly" + err.Error())
	}

	app.DbHost = os.Getenv("DB_HOST")
	app.DbPort = os.Getenv("DB_PORT")
	app.DbUser = os.Getenv("DB_USER")
	app.DbPassword = os.Getenv("DB_PASSWORD")
	app.DbName = os.Getenv("DB_NAME")
}
