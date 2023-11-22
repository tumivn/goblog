package cms

import (
	"github.com/labstack/echo/v4"
	"github.com/legangs/cms/internal/domain/cms/models"
	"github.com/legangs/cms/internal/domain/cms/repositories"
	"net/http"
	"time"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}
