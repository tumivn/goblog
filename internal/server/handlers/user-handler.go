package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/legangs/cms/internal/domain/auth/dtos"
	"github.com/legangs/cms/internal/domain/auth/services"
	"github.com/legangs/cms/internal/server"
	"net/http"
	"strings"
)

type UserHandler struct {
	server *server.Server
}

func CreateUserHandler(s *server.Server) *UserHandler {
	return &UserHandler{
		server: s,
	}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	u := dtos.CreateUserRequest{}
	c.Bind(&u)

	err := u.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser, err := services.CreatUser(u)

	if err != nil && strings.Contains(err.Error(), "internal error") {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}
