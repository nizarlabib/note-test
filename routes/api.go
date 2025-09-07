package routes

import (
	"sidita-be/controllers"
	"sidita-be/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	
	r.Use(middlewares.CORSMiddleware())

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

		api.GET("/auth/me", controllers.CurrentUser)

		user := api.Group("/user")
		{
			user.GET("/get", controllers.GetUsers)
			user.GET("/:id", controllers.GetUserByID)
		}

		note := api.Group("/note")
		{
			note.POST("/add", controllers.CreateNote)
			note.GET("/get", controllers.GetAllNote)
			note.GET("/get/:id", controllers.GetNoteByID)
			note.PUT("/update/:id", controllers.UpdateNote)
			note.DELETE("/delete/:id", controllers.DeleteNote)
		}
	}
}