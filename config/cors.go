package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS() *gin.Engine {
	corsRoute := gin.Default()

	corsRoute.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	return corsRoute
}
