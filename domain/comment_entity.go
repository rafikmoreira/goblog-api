package domain

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body string
}
