package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDBConnection(migrate func(DBConnection *gorm.DB)) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to migrate database")
	}
	migrate(db)
	return db
}
