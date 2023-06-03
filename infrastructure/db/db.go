package db

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate(DBConnection *gorm.DB) {
	err := DBConnection.AutoMigrate(
		&domain.Post{},
		&domain.User{},
		&domain.Comment{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}

func newDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to migrate database")
	}
	migrate(db)
	return db
}

var DBConnection = newDBConnection()
