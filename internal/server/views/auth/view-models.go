package auth

import "time"

type SignInViewModel struct {
	ErrorMessage string
	Email        string
}

type SignUpViewModel struct {
	ErrorMessage string
	Email        string
	Password     string
	Username     string
	Firstname    string
	Lastname     string
}

type UserProfileViewModel struct {
	Username  string
	Firstname string
	Lastname  string
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
