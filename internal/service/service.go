package service

import (
	"github.com/rommomm123321/go-rest-api/internal/entities"
	"github.com/rommomm123321/go-rest-api/internal/repository"
)

type Service struct {
	Dog
}
type Dog interface {
	Create(data entities.Dog) (entities.Dog, error)
	Update(data entities.Dog) (entities.Dog, error)
	Delete(dogID uint) error
	GetByID(dogID uint) (entities.Dog, error)
	GetAll() ([]entities.Dog, error)
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Dog: NewDogService(r.Dog),
	}
}
