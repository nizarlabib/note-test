package routes

import (
	"sidita-be/controllers"
	"sidita-be/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		api.Use(middlewares.JwtAuthMiddleware())

		api.GET("/me", controllers.CurrentUser)
		api.GET("/project/get", controllers.GetProjects)
		api.GET("/project/:id", controllers.GetProjectByID)
		api.GET("/worklog/get", controllers.GetWorklogs)
		api.GET("/worklog/:uid", controllers.GetWorklogsByUserID)
	}
}