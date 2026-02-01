package repositories

import (
	"context"

	"github.com/TaushifReza/go-event-booking-api/models"
	"gorm.io/gorm"
)

type EventRegistrationRepository struct {
	DB *gorm.DB
}

func NewEventRegistration(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (r *EventRegistrationRepository) Create(ctx context.Context, eventRegister *models.EventRegistration) error {
	return r.DB.WithContext(ctx).Create(eventRegister).Error
}

func (r *EventRegistrationRepository) IsAlreadyRegister(ctx context.Context, eventID, userID uint) (bool, error) {
	var exists int

	if err := r.DB.WithContext(ctx).Model(&models.EventRegistration{}).Select("1").Where("user_id = ? AND event_id = ?", userID, eventID).Limit(1).Scan(&exists).Error; err != nil {
		return false, err
	}

	return exists == 1, nil
}
