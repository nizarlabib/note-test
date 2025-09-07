package main

import (
	"log"
	"note-test/config"
	"note-test/models"
	"note-test/routes"

	_ "note-test/docs"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sidita Tes API
// @version 1.0
// @description API documentation for Sidita Tes

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// connect DB → assign ke global config.DB
	config.ConnectDB()

	log.Println("Migrating...")

	// pakai global config.DB
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Note{},
		&models.Log{},
	); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migration completed")

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.ApiRoutes(r)

	r.Run(":8080")
}

