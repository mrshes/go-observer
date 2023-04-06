package migration

import (
	"first-project/internal/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
