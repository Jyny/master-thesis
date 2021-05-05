package sql

import (
	"fmt"
	"server/pkg/model"
	"strings"

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

	// create enmu type
	createEnum(db, model.StatusType, []string{
		string(model.PENDING),
		string(model.RUNNING),
		string(model.COMPLETE),
	})

	// create enmu type
	createEnum(db, model.WorkerClassType, []string{
		string(model.ALIGN),
		string(model.ANC),
	})

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

func createEnum(db *gorm.DB, enumname string, enumvalues []string) {
	db.Exec(fmt.Sprintf(
		"CREATE TYPE %s AS ENUM (%s);",
		enumname,
		"'"+strings.Join(enumvalues, "', '")+"'",
	))
}
