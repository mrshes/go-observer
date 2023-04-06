package repository

import (
	"first-project/internal/database/models"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) Users {
	return Users{db: db}
}

func (u *Users) Create(user *models.User) (*models.User, error) {
	db := u.db.Where(user).FirstOrCreate(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return user, nil
}

func (u *Users) Get(user *models.User) {

}
