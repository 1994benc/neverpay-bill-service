package main

import (
	"1994benc/neverpay-api/internal/bill"
	"1994benc/neverpay-api/internal/database"
	transportHTTP "1994benc/neverpay-api/internal/transport/http"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type App struct{}

// Run - runs our application. We set it up in a struct so that it's easy for testing
func (app *App) Run() error {
	log.Println("Running the server")
	var err error
	var db *gorm.DB
	db, err = database.New()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	err = database.MigrateDB(db)
	if err != nil {
		log.Fatalf("Error migrating DB: %s", err)
	}
	billService := bill.New(db)
	handler := transportHTTP.New(billService)
	handler.SetupRoutes()
	err = http.ListenAndServe(":8080", handler.Router)
	return err
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app := App{}
	err := app.Run()
	if err != nil {
		log.Fatalf("Error starting the server %s", err)
	}
}
