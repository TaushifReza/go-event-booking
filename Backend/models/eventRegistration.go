package models

import "github.com/TaushifReza/go-event-booking-api/db"

type EventRegistration struct {
	ID      int64
	EventID int64
	UserID  int64
}

func RegisterForEvent(eventID, userID int64) (error){
	query := `INSERT INTO eventRegistrations(event_id, user_id) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(eventID, userID)
	if err != nil{
		return nil
	}
	return err
}