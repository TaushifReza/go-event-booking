package models

import (
	"time"

	"github.com/TaushifReza/go-event-booking-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID		int
}

var events = []Event{}

func (e Event) Save() error{
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil{
		return nil
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() []Event{
	return  events
}