package routes

import (
	"account-management/backend/controllers"
	"account-management/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/", middlewares.AuthMiddleware(), controllers.CreateUser)
		user.GET("/", middlewares.AuthMiddleware(), controllers.FindUsers)
		user.GET("/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
		user.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
		user.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}

}
