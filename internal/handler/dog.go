package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rommomm123321/go-rest-api/internal/entities"
)

func (h *Handler) AddDog(c *fiber.Ctx) error {
	dog := new(entities.Dog)

	if err := c.BodyParser(dog); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, "Unable to parse request body", err.Error())
	}

	if err := validateCreateStruct(dog); err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			return newErrorResponse(c, fiber.StatusBadRequest, validationErr.Message, validationErr.ValidationErrors)
		}
		return newErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	createdDog, err := h.services.Dog.Create(*dog)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, "Unable to create dog", err.Error())
	}

	return newResponse(c, fiber.StatusCreated, createdDog)
}

func (h *Handler) GetDogByID(c *fiber.Ctx) error {
	dogID := c.Params("id")
	id, err := strconv.Atoi(dogID)
	if err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, "Invalid dog ID", "")
	}

	dog, err := h.services.Dog.GetByID(uint(id))
	if err != nil {
		return newErrorResponse(c, fiber.StatusNotFound, "Dog not found", "")
	}

	return newResponse(c, fiber.StatusOK, dog)
}

func (h *Handler) GetAllDogs(c *fiber.Ctx) error {
	dogs, err := h.services.Dog.GetAll()
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, "Unable to fetch dogs", err.Error())
	}

	return newResponse(c, fiber.StatusOK, dogs)
}

func (h *Handler) UpdateDog(c *fiber.Ctx) error {
	dogID := c.Params("id")
	id, err := strconv.Atoi(dogID)
	if err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, "Invalid dog ID", "")
	}

	dog := new(entities.Dog)
	if err := c.BodyParser(dog); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, "Unable to parse request body", err.Error())
	}

	dog.ID = uint(id)

	if err := validateUpdateStruct(dog); err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			return newErrorResponse(c, fiber.StatusBadRequest, validationErr.Message, validationErr.ValidationErrors)
		}
		return newErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	updatedDog, err := h.services.Dog.Update(*dog)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, "Unable to update dog", err.Error())
	}

	return newResponse(c, fiber.StatusOK, updatedDog)
}

func (h *Handler) DeleteDog(c *fiber.Ctx) error {
	dogID := c.Params("id")
	id, err := strconv.Atoi(dogID)
	if err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, "Invalid dog ID", "")
	}

	if err := h.services.Dog.Delete(uint(id)); err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, "Unable to delete dog", err.Error())
	}

	return newResponse(c, fiber.StatusNoContent, nil)
}
