package config

import (
	"fmt"

	"github.com/efydb/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	dir := RootDir()
	MkDirIfNotExists(dir)

	dbUri := fmt.Sprintf("%s/appdb.sqlite", dir)
	Database, err = gorm.Open(
		sqlite.Open(dbUri),
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
