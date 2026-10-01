package main

import (
	"account-management/backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World tes",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "8080"))
}
