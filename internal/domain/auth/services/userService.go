package services

import (
	"errors"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/models"
	"github.com/tumivn/goblog/internal/domain/auth/repositories"
	"github.com/tumivn/goblog/ultilities"
	"time"
)

func CreatUser(dto dtos.CreateUserRequest) (*dtos.UserResponse, error) {

	err := dto.Validate()
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username:  dto.Username,
		Email:     dto.Email,
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.SetPassword(dto.Password) // hash password then save to user.Password

	_, err = repositories.GetUserByEmail(user.Email)
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

	return &dtos.UserResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		Firstname: newUser.Firstname,
		Lastname:  newUser.Lastname,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func AuthenticateUser(request *dtos.LoginRequest, jwtSecret string) (*dtos.LoginResponse, error) {
	user, err := repositories.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.CheckPassword(request.Password) == false {
		return nil, errors.New("invalid email or password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	tokenString, err := ultilities.GenerateJwt(request.Email, jwtSecret, expirationTime)

	if err != nil {
		return nil, errors.New("unable to create token")
	}

	return &dtos.LoginResponse{
		Token:   tokenString,
		Email:   user.Email,
		Expires: expirationTime,
	}, nil
}

func GetUsers() ([]dtos.UserResponse, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByEmail(email string) (*dtos.UserResponse, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
