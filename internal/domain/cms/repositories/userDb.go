package repositories

import (
	"github.com/tumivn/goblog/internal/domain/cms/models"
)

func CreateUser(user models.User) (models.User, error) {
	return user, nil
}
