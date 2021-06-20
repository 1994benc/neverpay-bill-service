package database

import (
	"github.com/jinzhu/gorm"
)

// Migrate our database and create bill table
func MigrateDB(db *gorm.DB) error {
	models := []interface{}{}
	result := db.AutoMigrate(models...)
	return result.Error
}
