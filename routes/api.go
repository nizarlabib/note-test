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

		api.GET("/me", controllers.CurrentUser)

		user := api.Group("/user")
		{
			user.GET("/get", controllers.GetUsers)
			user.GET("/:id", controllers.GetUserByID)
			user.GET("/get/absent/recap", controllers.GetUserAbsentRecap)
			user.GET("/get/absent/recap/byuser/:id", controllers.GetUserAbsentRecapByUID)
		}

		project := api.Group("/project")
		{
			project.GET("/get", controllers.GetProjects)
			project.GET("/:id", controllers.GetProjectByID)
		}

		worklog := api.Group("/worklog")
		{
			worklog.POST("/add", controllers.InsertWorklog)
			worklog.GET("/get", controllers.GetWorklogs)
			worklog.GET("/get/user/:user_id", controllers.GetWorklogsByUserId)
			worklog.DELETE("/delete/:id", controllers.DeleteWorklog)
			worklog.GET("/get/user/byproject/:user_id", controllers.GetTotalUserHoursWorkedByProject)
		}
	}
}