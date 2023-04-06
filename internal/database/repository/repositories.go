package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	DB    *gorm.DB
	Users Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		DB:    db,
		Users: NewUsers(db),
	}
}
