package main

import (
	"account-management/backend/config"
	"account-management/backend/database"
	"account-management/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	router := gin.Default()

	config.SetupCORS()
	routes.AuthRoute(router)
	routes.UserRoute(router)

	port := config.GetEnv("APP_PORT", "8080")
	router.Run(":" + port)

}
