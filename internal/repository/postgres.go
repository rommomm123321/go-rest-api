package repository

import (
	"github.com/rommomm123321/go-rest-api/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true})
	if err != nil {
		return nil, err
	}
	Database.AutoMigrate(&entities.Dog{})

	return Database, nil
}
