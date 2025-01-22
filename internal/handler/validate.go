package handler

import (
	"github.com/go-playground/validator/v10"
)

func validateCreateStruct(entity interface{}) error {
	validate := validator.New()

	err := validate.Struct(entity)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errorMessages[fieldErr.Field()] = fieldErr.Tag()
		}
		return &ValidationError{
			Message:          "Validation failed for creation",
			ValidationErrors: errorMessages,
		}
	}

	return nil
}

func validateUpdateStruct(entity interface{}) error {
	validate := validator.New()

	validate.RegisterValidation("optional_required", func(fl validator.FieldLevel) bool {
		return fl.Field().String() != "" || fl.Parent().FieldByName("ID").IsValid()
	})

	err := validate.Struct(entity)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errorMessages[fieldErr.Field()] = fieldErr.Tag()
		}
		return &ValidationError{
			Message:          "Validation failed for update",
			ValidationErrors: errorMessages,
		}
	}

	return nil
}

type ValidationError struct {
	Message          string            `json:"message"`
	ValidationErrors map[string]string `json:"validationErrors"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
