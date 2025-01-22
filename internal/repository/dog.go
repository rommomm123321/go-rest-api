package repository

import (
	"github.com/rommomm123321/go-rest-api/internal/entities"
	"gorm.io/gorm"
)

type DogRepository struct {
	db *gorm.DB
}

func NewDogRepository(db *gorm.DB) *DogRepository {
	return &DogRepository{db: db}
}

func (r *DogRepository) Create(dog entities.Dog) (entities.Dog, error) {
	if err := r.db.Create(&dog).Error; err != nil {
		return dog, err
	}
	return dog, nil
}

func (r *DogRepository) Update(dog entities.Dog) (entities.Dog, error) {
	if err := r.db.Save(&dog).Error; err != nil {
		return dog, err
	}
	return dog, nil
}

func (r *DogRepository) Delete(dogID uint) error {
	if err := r.db.Delete(&entities.Dog{}, dogID).Error; err != nil {
		return err
	}
	return nil
}

func (r *DogRepository) GetByID(dogID uint) (entities.Dog, error) {
	var dog entities.Dog
	if err := r.db.First(&dog, dogID).Error; err != nil {
		return dog, err
	}
	return dog, nil
}

func (r *DogRepository) GetAll() ([]entities.Dog, error) {
	var dogs []entities.Dog
	if err := r.db.Find(&dogs).Error; err != nil {
		return nil, err
	}
	return dogs, nil
}
