package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/internal/common"
	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/TaushifReza/go-event-booking-api/repositories"
)

type EventRegistrationService struct {
	Repo *repositories.EventRegistrationRepository
}

func NewEventRegistration(repo *repositories.EventRegistrationRepository) *EventRegistrationService {
	return &EventRegistrationService{Repo: repo}
}

func (s *EventRegistrationService) Register(ctx context.Context, eventID, userID uint) error {
	eventRegisteration := &models.EventRegistration{
		UserID:  userID,
		EventID: eventID,
	}

	// check event exist

	isAlreadyRegister, err := s.Repo.IsAlreadyRegister(ctx, eventID, userID)

	if err != nil {
		return &common.AppError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again later"),
		}
	}

	if isAlreadyRegister {
		return &common.AppError{
			Code:    http.StatusBadRequest,
			Message: fmt.Errorf("you have already register for this event"),
		}
	}

	if err := s.Repo.Create(ctx, eventRegisteration); err != nil {
		return &common.AppError{
			Code:    http.StatusBadRequest,
			Message: fmt.Errorf("something went wrong. please try again later"),
		}
	}

	return nil

}
