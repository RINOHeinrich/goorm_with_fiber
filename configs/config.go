package config

import (
	"github.com/RINOHeinrich/goorm_with_fiber/models"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		/* 		SkipDefaultTransaction: true,
		   		PrepareStmt:            true, */
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.Product{})

	return nil
}
