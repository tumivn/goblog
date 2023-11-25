package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/legangs/cms/internal/domain/cms/dtos"
	"github.com/legangs/cms/internal/domain/cms/models"
	"github.com/legangs/cms/internal/domain/cms/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func CreatUser(request dtos.CreateUserRequest) (*dtos.UserResponse, error) {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("unable to hash password: %v", err)
		return nil, errors.New("unable to create user")
	}
	user.Password = string(hashedPassword)

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

func AuthenticateUser(request dtos.LoginRequest, jwtSecret string) (*dtos.LoginResponse, error) {
	user, err := repositories.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &dtos.Claims{
		Email: request.Email,
		Claims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

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
