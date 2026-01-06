package errors

import (
	"errors"
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/internal/common"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, message string, err error){
	var appErr *common.AppError
	if errors.As(err, &appErr){
		ctx.JSON(appErr.Code, utils.ErrorResponse(message, appErr.Message))
		return
	}

	ctx.JSON(
		http.StatusInternalServerError, utils.ErrorResponse(
			"something went wrong. please try again", 
			errors.New("something went wrong. please try again")))
}