package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"spotsync-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get DATABASE_URL from .env
	dsn := os.Getenv("DATABASE_URL")

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
	err = DB.AutoMigrate(&models.User{}, &models.ParkingZone{}, &models.Reservation{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	log.Println("Database connected successfully")
}
