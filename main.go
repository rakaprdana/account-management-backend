package main

import (
	"account-management/backend/config"
	"account-management/backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World tes",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "8080"))
}
