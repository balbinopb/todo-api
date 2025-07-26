package database

import (
    "fmt"
    "log"
    "os"
    "todo-api/models"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

    // Load .env file
    _ = godotenv.Load()

	// Hlo/hamosu string koneksaun husi .env
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_TIMEZONE"),
    )

    var err error
	// Halo koneksaun ho database PostgreSQL  uza GORM
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }


	// 
    fmt.Println("Database connected!")

    // Auto halo migrasaun ba model Todo (se tabela la iha, GORM sei kria)
    DB.AutoMigrate(&models.Todo{})
}
