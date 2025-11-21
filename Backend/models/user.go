package models

import (
	"github.com/TaushifReza/go-event-booking-api/db"
	"github.com/TaushifReza/go-event-booking-api/utils"
)

type User struct {
	ID       int64 `json:"id"`
	Email    string `binding:"required" json:"email"`
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

func GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT * FROM
		users
		WHERE email = ? 
	`

	row := db.DB.QueryRow(query, email)

	var user User

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}