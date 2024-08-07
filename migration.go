package main

import (
    "fmt"
    "log" // tambahkan impor log
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "os"
    "payment-go/models" // impor package models
)

func init() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func main() {
    // Initialize database connection
    dsn := os.Getenv("DATABASE_URL") // Format: "host=localhost user=postgres dbname=payment_gateway sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Run migrations
    err = db.AutoMigrate(&models.User{}, &models.Coins{}, &models.Order{}, &models.Invoice{}, &models.TopUpTransaction{})
    if err != nil {
        panic("failed to migrate database: " + err.Error())
    }

    fmt.Println("Database migrated successfully!")
}
