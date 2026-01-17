package repositories

import "gorm.io/gorm"

type EventRegistrationRepository struct {
	DB *gorm.DB
}

func NewEventRegistration(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}
