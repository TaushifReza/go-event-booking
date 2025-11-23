package utils

import "errors"

func SuccessResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(message string, err error) map[string]interface{} {
	if err == nil {
		err = errors.New("unknown error")
	}
	return map[string]interface{}{
		"success": false,
		"message": message,
		"error":   err.Error(),
	}
}

func ValidationErrorResponse(errors map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"errors":  errors,
	}
}