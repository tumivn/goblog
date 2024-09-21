package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/internal/server"
	"github.com/tumivn/goblog/internal/storage"
)

func main() {
	s := server.NewServer()
	s.Init()
	db := storage.GetDB()

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/storage/migrations",
		"postgres", driver)

	if err != nil {
		panic(err)
	}
	err = m.Down()
	if err != nil {
		panic(err)
	}
}
