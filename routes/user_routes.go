package routes

import (
	"go-user-psql/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/users")

	{
		user.POST("/register", controllers.CreateUser)
		user.POST("/login", controllers.Login)
		user.GET("/", controllers.GetUsers)
	}
}
