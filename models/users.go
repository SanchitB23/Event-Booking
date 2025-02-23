package models

import (
	"errors"
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

func (user *User) Authenticate() error {
	var hashedPassword string
	err := db.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Email).Scan(&hashedPassword)
	if err != nil {
		return errors.New("user not found")
	}
	if !utils.VerifyPassword(user.Password, hashedPassword) {
		return errors.New("invalid password")
	}
	return nil
}
