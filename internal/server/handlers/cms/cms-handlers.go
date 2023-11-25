package cms

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/legangs/cms/internal/domain/cms/dtos"
	"github.com/legangs/cms/internal/domain/cms/services"
	"github.com/legangs/cms/internal/server"
	"net/http"
	"time"
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

	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   res.Token,
		Expires: res.Expires,
	})

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetMe(c echo.Context) error {

	token, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	t, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.server.Config.JwtSecret), nil
	})

	if err != nil || !t.Valid {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	user, err := services.GetUserByEmail(t.Claims.(jwt.MapClaims)["email"].(string))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Logout(c echo.Context) error {
	_, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})

	return c.JSON(
		http.StatusOK,
		&dtos.MessageResponse{
			Message: "Logout success",
		},
	)
}
