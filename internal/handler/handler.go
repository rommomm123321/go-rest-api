package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rommomm123321/go-rest-api/internal/service"
)

type Handler struct {
	services  *service.Service
	validator *validator.Validate
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services:  services,
		validator: validator.New(),
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	api := router.Group("/api")
	api.Post("/dogs", h.AddDog)
	api.Get("/dogs", h.GetAllDogs)
	api.Get("/dogs/:id", h.GetDogByID)
	api.Put("/dogs/:id", h.UpdateDog)
	api.Delete("/dogs/:id", h.DeleteDog)

	return router
}
