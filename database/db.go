package database

import (
	"fmt"
	"log"
	"os"

	"organization-service/models"

	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

var DB *gorm.DB
var TestDB *gorm.DB

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

	database.AutoMigrate(&models.Organization{}, &models.Member{}, &models.OrganizationMember{})
	DB = database
}
func ConnectTestDB() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME") + "_test"
	fmt.Print("Database name: ", databaseName)
	fmt.Print("Host: ", host)
	fmt.Print("Username: ", username)
	fmt.Print("Password: ", password)
	connectionString := "sqlserver://" + username + ":" + password + "@" + host + "?database=" + databaseName
	testDatabase, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to test database:", err)
	}

	testDatabase.AutoMigrate(&models.Organization{}, &models.Member{}, &models.OrganizationMember{})
	TestDB = testDatabase
	fmt.Println("Test database connected")
}
