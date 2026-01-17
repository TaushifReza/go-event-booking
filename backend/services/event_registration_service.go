package services

import "github.com/TaushifReza/go-event-booking-api/repositories"

type EventRegistrationService struct {
	Repo *repositories.EventRegistrationRepository
}

func NewEventRegistration(repo *repositories.EventRegistrationRepository) *EventRegistrationService {
	return &EventRegistrationService{Repo: repo}
}
