package database

import (
	"1994benc/neverpay-api/internal/user"

	"github.com/jinzhu/gorm"
)

// Migrate our database and create bill table
func MigrateDB(db *gorm.DB) error {
	models := []interface{}{&user.User{}}
	result := db.AutoMigrate(models...)
	return result.Error
}
