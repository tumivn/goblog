package home

import "github.com/tumivn/goblog/internal/domain/auth/dtos"

type ViewModel struct {
	ErrorMessage string
	User         *dtos.UserResponse
}
