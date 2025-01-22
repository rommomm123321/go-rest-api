package app

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, app *fiber.App) error {
	return app.Listen(":" + port)
}
