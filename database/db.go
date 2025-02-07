package database

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"organization-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB
var TestDB *gorm.DB

func loadEnv(isTestDB bool) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	var path string
	if isTestDB {
		path = filepath.Join(filepath.Dir(dir), ".env")
	} else {
		path = filepath.Join(dir, "/.env")
	}

	err = godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getConnectionString(databaseName string) string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	encodedUsername := url.QueryEscape(username)
	encodedPassword := url.QueryEscape(password)

	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", encodedUsername, encodedPassword, host, databaseName)
}

func connectDatabase(databaseName string) (*gorm.DB, error) {
	connectionString := getConnectionString(databaseName)
	return gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
}

func ConnectDatabase() {
	loadEnv(false)
	databaseName := os.Getenv("DB_NAME")
	database, err := connectDatabase(databaseName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	database.AutoMigrate(&models.Organization{}, &models.Member{}, &models.OrganizationMember{})
	DB = database
}

func ConnectTestDB() {
	loadEnv(true)
	databaseName := os.Getenv("DB_NAME") + "_test"
	testDatabase, err := connectDatabase(databaseName)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}
	testDatabase.AutoMigrate(&models.Organization{}, &models.Member{}, &models.OrganizationMember{})
	TestDB = testDatabase
	DB = testDatabase
	fmt.Println("Test database connected")
}

func MigrateTestDatabase() {
	TestDB.AutoMigrate(&models.Organization{}, &models.Member{}, &models.OrganizationMember{})
}
