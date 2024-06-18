package config

import (
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	//Database.AutoMigrate(&entities.Dog{})

	return nil
}
