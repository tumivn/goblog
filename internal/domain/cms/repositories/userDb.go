package repositories

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/legangs/cms/internal/domain/cms/dtos"
	"github.com/legangs/cms/internal/domain/cms/models"
	"github.com/legangs/cms/internal/storage"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (username, email, firstName, lastName, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6,$7) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Username, user.Email, user.Firstname, user.Lastname, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM users WHERE username=$1`
	var user models.User
	err := db.QueryRow(sqlStatement, username).Scan(&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM users WHERE email=$1`
	var user models.User
	err := db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUsers() ([]dtos.UserResponse, error) {
	db := storage.GetDB()
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Select("id", "username", "email", "firstname", "lastname", "created_at", "updated_at").
		From("users")

	rows, err := query.RunWith(db).Query()

	if err != nil {
		return nil, err
	}

	var users []dtos.UserResponse
	for rows.Next() {
		var user dtos.UserResponse
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
