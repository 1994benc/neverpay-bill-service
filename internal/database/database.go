package database

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Create a new database connection
func New() (*gorm.DB, error) {
	log.Println("Setting up new database connection")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	sslMode := os.Getenv("SSL_MODE")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPassword, sslMode)
	log.Printf("Connecting to database with the following credentials: %s", connectionStr)
	db, err := gorm.Open("postgres", connectionStr)
	return db, err
}
