package main

import (
	"account-management/backend/config"
	"account-management/backend/database"
	"account-management/backend/routes"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	userRoute := routes.UserRoute()

	userRoute.Run(":" + config.GetEnv("APP_PORT", "8080"))
}
