package db

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"gorm.io/gorm"
)

func Migrate(DBConnection *gorm.DB) {
	err := DBConnection.AutoMigrate(
		&domain.Post{},
		&domain.User{},
		&domain.Comment{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
