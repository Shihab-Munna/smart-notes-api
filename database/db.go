package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

// Init initializes the database connection using GORM
func Init() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN environment variable not set")
    }


    var err error

    // Open the PostgreSQL database connection using GORM
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Ping the database to verify the connection
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to extract database connection: %v", err)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }

    log.Println("Connected to PostgreSQL database successfully!")
}
