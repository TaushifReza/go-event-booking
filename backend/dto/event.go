package dto

import "time"

type EventCreateDto struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Venue       string    `json:"venue" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
}

type EventDetailResponseDto struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Venue       string    `json:"venue"`
	DateTime    time.Time `json:"date_time"`
}
