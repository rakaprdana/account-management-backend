package routes

import (
	"account-management/backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		// Authentication
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

}
