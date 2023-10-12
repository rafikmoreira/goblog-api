package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLConnection(migrate func(DBConnection *gorm.DB)) *gorm.DB {
	dsn := "host=localhost user=postgre password=postgre dbname=postgre port=5432 sslmode=disable TimeZone=America/St_Kitts"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to migrate database")
	}
	migrate(db)
	return db
}
