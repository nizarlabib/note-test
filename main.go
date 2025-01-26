package main

import (
	"sidita-be/config"
	"sidita-be/routes"

	_ "sidita-be/docs"

	"github.com/gin-gonic/gin"
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

	config.ConnectDB()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.ApiRoutes(r)

	r.Run(":8080")
}
