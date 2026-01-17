package db

import "github.com/TaushifReza/go-event-booking-api/models"

func Migrate() error {
	return DB.AutoMigrate(&models.User{}, &models.Event{}, &models.EventRegistration{})
}
