package test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnect(model *interface{}) *gorm.DB {
	dsn := "file::memory:"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(model)
	if err != nil {
		panic(err)
	}

	return db
}
