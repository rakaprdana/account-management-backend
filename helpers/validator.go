package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorMap := make(map[string]string)

	if validationError, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationError {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorMap[field] = fmt.Sprintf(" %s is required", field)
			case "email":
				errorMap[field] = "Invalid email format"
			case "unique":
				errorMap[field] = fmt.Sprintf("%s already exist", field)
			case "min":
				errorMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				errorMap[field] = fmt.Sprint("%s must be a number", field)
			default:
				errorMap[field] = "Invalid value"
			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorMap["Username"] = "Username already exist"
			}

			if strings.Contains(err.Error(), "email") {
				errorMap["Email"] = "Email already exist"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorMap["Error"] = "Record not found"
		}
	}

	return errorMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
