package migrations

import (
	"github.com/YagoAyala/go-rest-api.git/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}
