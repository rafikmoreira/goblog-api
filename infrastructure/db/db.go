package db

import (
	"github.com/rafikmoreira/go-blog-api/domain"
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

var Connection = NewSQLiteDBConnection(migrate)
