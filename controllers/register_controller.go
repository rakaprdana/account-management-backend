package controllers

import (
	"account-management/backend/database"
	"account-management/backend/helpers"
	"account-management/backend/models"
	"account-management/backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Error validation",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name:      req.Name,
		Username:  req.Username,
		Email:     req.Email,
		IsDeleted: false,
		Password:  helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {

			// Error duplicate entry
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			// Another error (500 Response)
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}

		return
	}
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			IsDeleted: user.IsDeleted,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:29"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:29"),
		},
	})
}
