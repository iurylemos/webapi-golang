package migrations

import (
	"github.com/iurylemos/webapi-golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}