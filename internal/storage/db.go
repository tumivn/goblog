package storage

import (
	"database/sql"
	"fmt"
	"github.com/tumivn/goblog/internal/server/config"
)

var db *sql.DB

func InitDB(app *config.AppConfig) {
	var err error
	db, err = sql.Open("postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			app.DbHost, app.DbUser, app.DbPassword, app.DbName, app.DbPort))

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}
