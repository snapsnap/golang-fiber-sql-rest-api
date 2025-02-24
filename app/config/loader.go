package config

import (
	"log"
	"os"
	"rest-project/app/models"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *models.Configuration {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error when loading file configuration : ", err.Error())
	}

	expInt, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	return &models.Configuration{
		Server: models.Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: models.Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Tz:   os.Getenv("DB_TIMEZONE"),
		},
		Jwt: models.Jwt{
			Key: os.Getenv("JWT_SECRET"),
			Exp: expInt,
		},
	}
}
