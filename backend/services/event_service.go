package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/internal/common"
	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/TaushifReza/go-event-booking-api/repositories"
	"gorm.io/gorm"
)

type EventService struct {
	Repo *repositories.EventRepository
}

func NewEventService(repo *repositories.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (e *EventService) Create(ctx context.Context, reqDto *dto.EventCreateDto, userID uint) (*models.Event, error) {
	event := &models.Event{
		Name:        reqDto.Name,
		Description: reqDto.Description,
		Location:    reqDto.Location,
		Venue:       reqDto.Venue,
		DateTime:    reqDto.DateTime,
		UserID:      &userID,
	}

	if err := e.Repo.Create(ctx, event); err != nil {
		return nil, &common.AppError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again"),
		}
	}
	return event, nil
}

func (e *EventService) GetAll(ctx context.Context, page, limit int) (*dto.PaginatedResponse[models.Event], error) {
	events, total, err := e.Repo.GetAll(ctx, page, limit)

	if err != nil {
		return nil, &common.AppError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again"),
		}
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.PaginatedResponse[models.Event]{
		Data:       events,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (e *EventService) GetByID(ctx context.Context, eventID uint) (*dto.EventDetailResponseDto, error) {
	event, err := e.Repo.GetByID(ctx, eventID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &common.AppError{
				Code:    http.StatusNotFound,
				Message: fmt.Errorf("event not found"),
			}
		}

		return nil, &common.AppError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again."),
		}
	}

	return &dto.EventDetailResponseDto{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		Venue:       event.Venue,
		DateTime:    event.DateTime,
	}, nil
}

func (e *EventService) Update(ctx context.Context, id, userID uint, dto *dto.EventUpdateDto) error {
	if !dto.HasUpdates() {
		return &common.AppError{
			Code:    http.StatusBadRequest,
			Message: errors.New("no fields provided to update"),
		}
	}

	if err := e.Repo.Update(ctx, id, userID, dto); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &common.AppError{
				Code:    http.StatusBadRequest,
				Message: fmt.Errorf("event with id: %v not found", id),
			}
		}

		return &common.AppError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again later"),
		}
	}

	return nil
}
