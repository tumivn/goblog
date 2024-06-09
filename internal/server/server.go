package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/internal/server/config"
	"github.com/tumivn/goblog/internal/storage"
)

type Server struct {
	Version   string
	Config    config.AppConfig
	validator *validator.Validate
	Echo      *echo.Echo
}

func NewServer() *Server {
	return &Server{
		Version:   "0.0.1",
		Config:    config.AppConfig{},
		validator: validator.New(),
		Echo:      echo.New(),
	}
}

func (s *Server) Init() {
	config.LoadConfig(&s.Config)
	storage.InitDB(&s.Config)
}

func (s *Server) Start(addr string) {
	s.Echo.Logger.Fatal(s.Echo.Start(":" + addr))
}
