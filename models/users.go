package models

import (
	"max-tuts/event-booking-rest-api/db"
	"max-tuts/event-booking-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user *User) Save() error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, hashedPassword)
	return err
}
