package config

import (
	"fmt"
	"log"

	// "os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env file
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get environment variables
	dbHost :=  env["DB_HOST"]
	dbPort :=  env["DB_PORT"]
	dbUser :=  env["DB_USER"]
	dbPassword :=  env["DB_PASSWORD"]
	dbName :=  env["DB_NAME"]
	// dbHost := "103.175.221.93"
	// dbPort := "5431"
	// dbUser := "postgres"
	// dbPassword := "merdeka45"
	// dbName := "sidita_nizar"

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatalf("Missing required environment variables for database connection")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	log.Println("Connecting to database with DSN:", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully!")
}
