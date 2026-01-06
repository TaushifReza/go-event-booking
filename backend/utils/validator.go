package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			field := fieldErr.Field()
			tag := fieldErr.Tag()

			switch tag {
			case "required":
				errors[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errors[field] = "Invalid email format"
			case "min":
				errors[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldErr.Param())
			default:
				errors[field] = fmt.Sprintf("%s is invalid", field)
			}
		}
	}
	return errors
}
