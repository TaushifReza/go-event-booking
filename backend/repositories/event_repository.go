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
func (r *EventRepository) GetAll(ctx context.Context, page, limit int) ([]models.Event, int64, error) {
	// declare variable
	var (
		events []models.Event
		total  int64
		offset = (page - 1) * limit
	)

	// count first
	if err := r.DB.
		WithContext(ctx).
		Model(&models.Event{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	// fetch page
	if err := r.DB.
		WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&events).
		Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

// get by id
func (r *EventRepository) GetByID(ctx context.Context, eventID uint) (*models.Event, error) {
	var event models.Event

	if err := r.DB.WithContext(ctx).First(&event, eventID).Error; err != nil {
		return nil, err
	}

	return &event, nil
}
