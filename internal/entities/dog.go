package entities

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string `json:"name" validate:"required,min=2"`
	Breed     string `json:"breed" validate:"required"`
	Age       int    `json:"age" validate:"gte=0,lte=25"`
	IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
	Color     string `json:"color" validate:"omitempty"`
}
