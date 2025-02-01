package database

import (
	"log"
	"os"

	"organization-service/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME")

	connectionString := "sqlserver://" + username + ":" + password + "@" + host + "?database=" + databaseName
	database, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	database.AutoMigrate(&models.Organization{})

	DB = database
}
