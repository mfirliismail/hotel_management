package utils

import (
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "hotel-management/models"
)

func InitDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    
    db.AutoMigrate(&models.User{}, &models.Room{}, &models.Booking{}, &models.Invoice{}, &models.Payment{})

    return db
}
