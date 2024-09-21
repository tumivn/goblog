package main

import (
	"github.com/bxcodec/faker/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/tumivn/goblog/internal/domain/auth/dtos"
	"github.com/tumivn/goblog/internal/domain/auth/services"
	"github.com/tumivn/goblog/internal/server"
)

func main() {
	s := server.NewServer()
	s.Init()

	users := []dtos.CreateUserRequest{
		{
			Username:  "xuanthuy",
			Email:     "xuanthuy@mail.com",
			Firstname: "Xuan",
			Lastname:  "Thuy",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "thuyxuan",
			Email:     "thuyxuan@mail.com",
			Firstname: "Thuy",
			Lastname:  "Xuan",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "ngocmai",
			Email:     "ngocmai@mail.com",
			Firstname: "Ngoc",
			Lastname:  "Mai",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "maingoc",
			Email:     "maingoc@mail.com",
			Firstname: "Mai",
			Lastname:  "Ngoc",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "thuyngoc",
			Email:     "thuyngoc@mail.com",
			Firstname: "Thuy",
			Lastname:  "Ngoc",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "ngocthuy",
			Email:     "ngocthuy@mail.com",
			Firstname: "Ngoc",
			Lastname:  "Thuy",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "thuyngocmai",
			Email:     "thuyngocmai@mail.com",
			Firstname: "Thuy",
			Lastname:  "Ngoc Mai",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "maingocmai",
			Email:     "mainm@mail.com",
			Firstname: "Mai",
			Lastname:  "Ngoc Mai",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "ngocle",
			Email:     "ngocle@mail.com",
			Firstname: "Ngoc",
			Lastname:  "Le",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "mainguyen",
			Email:     "mainguyen@mail.com",
			Firstname: "Mai",
			Lastname:  "Nguyen",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "hoangvanmai",
			Email:     "hoangvanmain@mail.com",
			Firstname: "Hoang",
			Lastname:  "Van Mai",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "maithich",
			Email:     "mainthich@mail.com",
			Firstname: "Mai",
			Lastname:  "Thich",
			Password:  "abc@ABC*123",
		},
		{
			Username:  "buunguyen",
			Email:     "buunguyen@mail.com",
			Firstname: "Buu",
			Lastname:  "Nguyen",
			Password:  "abc@ABC*123",
		},
	}

	for i := 0; i < 100; i++ {
		dto := dtos.CreateUserRequest{
			Username:  faker.Username(),
			Email:     faker.Email(),
			Firstname: faker.FirstName(),
			Lastname:  faker.LastName(),
			Password:  "abc@ABC*123",
		}
		_, err := services.CreatUser(dto)
		if err != nil {
			log.Error("Can't create user: ", err)
		}
	}

	for _, user := range users {
		u, _ := services.GetUserByEmail(user.Email)
		if u != nil {
			log.Infof("User ", u.Email, "is already existed")
		} else {
			_, err := services.CreatUser(user)
			if err != nil {
				log.Error("Can't create user: ", err)
			}
		}
	}
}
