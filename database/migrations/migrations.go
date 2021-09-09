package migrations

import (
	"github.com/YagoAyala/go-rest-api.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
