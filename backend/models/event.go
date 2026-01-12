package models

import (
	"time"
)

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Location    string    `gorm:"not null"`
	Venue       string    `gorm:"not null"`
	DateTime    time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      *uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
