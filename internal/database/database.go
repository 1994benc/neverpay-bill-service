package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Create a new database connection
func New() (*gorm.DB, error) {
	log.Println("Setting up new database connection")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbTable, dbPassword)
	log.Printf("Connecting to database with the following credentials: %s", connectionStr)
	db, err := gorm.Open("postgres", connectionStr)
	return db, err
}
