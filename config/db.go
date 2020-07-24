package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error

	username := os.Getenv("DB_USERNAME_PROD")
	password := os.Getenv("DB_PASSWORD_PROD")
	database := os.Getenv("DB_DATABASE_PROD")
	url := os.Getenv("DB_URL_PROD")
	port := os.Getenv("DB_PORT_PROD")

	if os.Getenv("DB_MODE") == "DEV" {
		username = os.Getenv("DB_USERNAME_DEV")
		password = os.Getenv("DB_PASSWORD_DEV")
		database = os.Getenv("DB_DATABASE_DEV")
		url = os.Getenv("DB_URL_DEV")
		port = os.Getenv("DB_PORT_DEV")
	}
	
	conection := username + ":" + password + "@tcp(" + url + ":" + port + ")/" + database

	DB, err = gorm.Open("mysql", conection)
	if err != nil {
		panic("failed to connect database")
	}
}
