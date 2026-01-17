package controllers

import (
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/gin-gonic/gin"
)

type EventRegistrationController struct {
	EventRegistrationService *services.EventRegistrationService
}

func NewEventRegistration(service *services.EventRegistrationService) *EventRegistrationController {
	return &EventRegistrationController{EventRegistrationService: service}
}

func (c *EventRegistrationController) Register(ctx *gin.Context) {}
