package main

import (
	"go-user-psql/config"
	"go-user-psql/models"
	"go-user-psql/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.Run(":8000")
}
