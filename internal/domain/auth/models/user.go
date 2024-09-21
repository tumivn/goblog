package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthdate time.Time `json:"birthdate"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
