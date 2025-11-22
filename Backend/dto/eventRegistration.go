package dto

type EventRegistrationCreate struct {
	EventID int64 `json:"event_id" binding:"required"`
	UserID  int64 `json:"user_id" binding:"required"`
}