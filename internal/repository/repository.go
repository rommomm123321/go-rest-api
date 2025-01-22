package repository

import (
	"github.com/rommomm123321/go-rest-api/internal/entities"
	"gorm.io/gorm"
)

type Repository struct {
	Dog
}

type Dog interface {
	Create(data entities.Dog) (entities.Dog, error)
	Update(data entities.Dog) (entities.Dog, error)
	Delete(dogID uint) error
	GetByID(dogID uint) (entities.Dog, error)
	GetAll() ([]entities.Dog, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Dog: NewDogRepository(db),
	}
}
