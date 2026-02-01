package controllers

import (
	"net/http"
	"strconv"

	"github.com/TaushifReza/go-event-booking-api/internal/http/errors"
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

type EventRegistrationController struct {
	EventRegistrationService *services.EventRegistrationService
}

func NewEventRegistration(service *services.EventRegistrationService) *EventRegistrationController {
	return &EventRegistrationController{EventRegistrationService: service}
}

func (e *EventRegistrationController) Register(ctx *gin.Context) {
	eventID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("invalid id", err))
		return
	}

	c := ctx.Request.Context()
	userID := ctx.GetUint("user_id")
	err = e.EventRegistrationService.Register(c, uint(eventID), userID)

	if err != nil {
		errors.HandleError(ctx, "failed to register for event", err)
		return
	}
    
    ctx.JSON(http.StatusCreated, utils.SuccessResponse("Successfuly register for event.", eventID))
}
