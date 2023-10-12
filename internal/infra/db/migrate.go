package db

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
	"gorm.io/gorm"
)

func Migrate(DBConnection *gorm.DB) {
	err := DBConnection.AutoMigrate(
		&entity.Post{},
		&entity.User{},
		&entity.Comment{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
