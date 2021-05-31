package database

import (
	"1994benc/neverpay-api/internal/bill"

	"github.com/jinzhu/gorm"
)

// Migrate our database and create bill table
func MigrateDB(db *gorm.DB) error {
	result := db.AutoMigrate(&bill.Bill{})
	return result.Error
}
