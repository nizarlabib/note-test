package handler

import (
	"net/http"
	"sync"

	"sidita-be/config"
	// "sidita-be/routes"

	_ "sidita-be/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	r    *gin.Engine
	once sync.Once
)

func setupRouter() {
}

func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
	config.ConnectDB()
	
	r = gin.Default()
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	// routes.ApiRoutes(r)

}
