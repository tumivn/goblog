package services

import (
	"errors"
	"github.com/legangs/cms/internal/domain/cms/dtos"
	"github.com/legangs/cms/internal/domain/cms/models"
	"github.com/legangs/cms/internal/domain/cms/repositories"
	"time"
)

func CreatUser(request dtos.CreateUserRequest) (*dtos.CreateUserResponse, error) {
	user := models.User{
		Username:  request.Username,
		Email:     request.Email,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := repositories.GetUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	_, err = repositories.GetUserByUsername(user.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	newUser, err := repositories.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return &dtos.CreateUserResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		Firstname: newUser.Firstname,
		Lastname:  newUser.Lastname,
	}, nil
}
