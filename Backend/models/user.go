package models

import (
	"github.com/TaushifReza/go-event-booking-api/db"
	"github.com/TaushifReza/go-event-booking-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error{
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil{
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}