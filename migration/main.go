package main

import (
	"log"
	"sidita-be/config"
	"sidita-be/models"
)

func main() {
	// connect ke DB
	config.ConnectDB()

	log.Println("Migrating...")

	// AutoMigrate return error langsung, bukan *gorm.DB
	if err := config.DB.AutoMigrate(
		// &models.User{},
		// &models.Note{},
		&models.Log{},
	); err != nil {
		log.Fatal("Migration failed: ", err)
	}


	log.Println("Migration completed")
}
