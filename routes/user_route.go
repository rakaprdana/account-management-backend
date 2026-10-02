package routes

import (
	"account-management/backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute() *gin.Engine {
	router := gin.Default()

	// Authentication
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	return router
}
