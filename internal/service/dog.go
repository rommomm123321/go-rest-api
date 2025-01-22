package service

import (
	"github.com/rommomm123321/go-rest-api/internal/entities"
	"github.com/rommomm123321/go-rest-api/internal/repository"
)

type DogService struct {
	repo repository.Dog
}

func NewDogService(repo repository.Dog) *DogService {
	return &DogService{repo: repo}
}

func (s *DogService) Create(dog entities.Dog) (entities.Dog, error) {
	return s.repo.Create(dog)
}

func (s *DogService) Update(dog entities.Dog) (entities.Dog, error) {
	return s.repo.Update(dog)
}

func (s *DogService) Delete(dogID uint) error {
	return s.repo.Delete(dogID)
}

func (s *DogService) GetByID(dogID uint) (entities.Dog, error) {
	return s.repo.GetByID(dogID)
}

func (s *DogService) GetAll() ([]entities.Dog, error) {
	return s.repo.GetAll()
}
