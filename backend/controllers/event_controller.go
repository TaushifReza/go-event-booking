package controllers

import (
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/internal/http/errors"
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{EventService: service}
}

func (e *EventController) Create(ctx *gin.Context) {
	var req dto.EventCreateDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.FormatValidationErrors(err)
		if len(validationError) > 0 {
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationError))
			return
		}

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid requst body", err))
		return
	}

	c := ctx.Request.Context()
	userID := ctx.GetUint("user_id")
	event, err := e.EventService.Create(c, &req, userID)
	if err != nil {
		errors.HandleError(ctx, "Event Creation failed.", err)
		return
	}

	res := dto.EventDetailResponseDto{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		Venue:       event.Venue,
		DateTime:    event.DateTime,
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Event created successfully", res))
}

func (e *EventController) GetAll(ctx *gin.Context) {
	c := ctx.Request.Context()
	events, err := e.EventService.GetAll(c)

	if err != nil {
		errors.HandleError(ctx, "Error fetching events.", err)
		return
	}

	result := make([]dto.EventDetailResponseDto, 0, len(events))

	for _, event := range events {
		result = append(result, dto.EventDetailResponseDto{ID: event.ID, Name: event.Name, Description: event.Description, Location: event.Location, Venue: event.Venue, DateTime: event.DateTime})
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("List of events.", result))
}
