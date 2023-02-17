package db

import (
	folder "changeme/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBase struct {
	SqlDB *gorm.DB
}

func (d DataBase) DbConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("organizacion.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(folder.Folder{})
	return db
}
