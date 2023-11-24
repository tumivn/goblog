package cms

import (
	"github.com/labstack/echo/v4"
	"github.com/legangs/cms/internal/domain/cms/dtos"
	"github.com/legangs/cms/internal/domain/cms/services"
	"github.com/legangs/cms/internal/server"
	"net/http"
)

type UserHandler struct {
	server *server.Server
}

func NewUserHandler(s *server.Server) *UserHandler {
	return &UserHandler{
		server: s,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	u := dtos.CreateUserRequest{}
	c.Bind(&u)

	err := u.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser, err := services.CreatUser(u)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) Login(c echo.Context) error {
	u := dtos.LoginRequest{}
	c.Bind(&u)

	err := u.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := services.AuthenticateUser(u, h.server.Config.JwtSecret)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
