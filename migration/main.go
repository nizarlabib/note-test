package main

import (
	"log"
	"note-test/config"
	"note-test/models"
)

func Migration() {
	// connect ke DB
	config.ConnectDB()

	log.Println("Migrating...")

	// AutoMigrate return error langsung, bukan *gorm.DB
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Note{},
		&models.Log{},
	); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migration completed")
}
