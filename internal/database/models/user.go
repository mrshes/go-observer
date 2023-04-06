package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
