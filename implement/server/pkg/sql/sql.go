package sql

import (
	"fmt"
	"server/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GromInit(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Migration(db)

	return db
}

func Migration(db *gorm.DB) {
	var err error

	// create extension
	createExtension(db, "uuid-ossp")

	// auto migration
	err = db.AutoMigrate(&model.Meeting{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Owner{})
	if err != nil {
		panic(err)
	}
}

func createExtension(db *gorm.DB, ext string) {
	db.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS \"%s\";", ext))
}
