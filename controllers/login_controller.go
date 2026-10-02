package controllers

import (
	"account-management/backend/database"
	"account-management/backend/helpers"
	"account-management/backend/models"
	"account-management/backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	// Search user based on username
	if err := database.DB.Where("username=?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	token := helpers.GeneratorToken(user.Username)
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token:     &token,
		},
	})
}
