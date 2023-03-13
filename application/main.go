package main

import (
	"envoyTokenGenerator/controllers"
	"envoyTokenGenerator/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// initializing database
	connectionString, port := SetENVs()
	database.Connect(connectionString)
	database.Migrate()

	// initialize router
	router := routerInit()
	router.Run(port)
}

func routerInit() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/token-management/token", controllers.GenerateToken)
		api.POST("/user-management/registration", controllers.RegisterUser)
	}
	return router
}
