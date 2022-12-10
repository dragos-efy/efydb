package config

import (
	"github.com/efydb/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

const DATABASE_URI = "postgres://postgres@localhost/efydb"

func Connect() error {
	var err error

	Database, err = gorm.Open(
		postgres.Open(DATABASE_URI),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.User{}, &entities.Theme{})

	return nil
}
