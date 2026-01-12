package repositories

import (
	"context"

	"github.com/TaushifReza/go-event-booking-api/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

// CREATE
func (r *EventRepository) Create(ctx context.Context, event *models.Event) error {
	return r.DB.WithContext(ctx).Create(event).Error
}

// Get all
func (r *EventRepository) GetAll(ctx context.Context) ([]models.Event, error) {
	var events []models.Event

	err := r.DB.WithContext(ctx).Find(&events).Error

	return events, err
}
