package controllers

import (
	"net/http"
	"strconv"

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
	var query dto.PaginationQueryDto
	if err := ctx.ShouldBindQuery(&query); err != nil {
		validationError := utils.FormatValidationErrors(err)
		if len(validationError) > 0 {
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationError))
			return
		}

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	query.Normalize()

	c := ctx.Request.Context()
	result, err := e.EventService.GetAll(c, query.Page, query.Limit)

	if err != nil {
		errors.HandleError(ctx, "Error fetching events.", err)
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("List of events.", result))
}

func (e *EventController) GetByID(ctx *gin.Context) {
	eventID, err := strconv.ParseUint(ctx.Param("eventID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("invalid id", err))
		return
	}

	c := ctx.Request.Context()

	result, err := e.EventService.GetByID(c, uint(eventID))

	if err != nil {
		errors.HandleError(ctx, "Error retriving event", err)
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Successfuly retrived event", result))
}

func (e *EventController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("invalid id", err))
		return
	}

	var req dto.EventUpdateDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.FormatValidationErrors(err)
		if len(validationError) > 0 {
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationError))
			return
		}
	}

	c := ctx.Request.Context()
	userID := ctx.GetUint("user_id")

	if err := e.EventService.Update(c, uint(id), userID, &req); err != nil {
		errors.HandleError(ctx, "error updating event", err)
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Event updated success", req))
}
