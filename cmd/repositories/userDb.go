package repositories

import "github.com/tumivn/goblog/cmd/models"

func CreateUser(user models.User) (models.User, error) {
	return user, nil
}
