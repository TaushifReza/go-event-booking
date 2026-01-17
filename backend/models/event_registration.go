package models

type EventRegistration struct {
	ID      uint  `gorm:"primaryKey"`
	UserID  uint  `gorm:"not null"`
	User    User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	EventID uint  `gorm:"not null"`
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
