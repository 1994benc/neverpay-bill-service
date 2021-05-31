package main

import (
	"1994benc/neverpay-api/internal/bill"
	"1994benc/neverpay-api/internal/database"
	transportHTTP "1994benc/neverpay-api/internal/transport/http"
	"net/http"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

// Run - runs our application. We set it up in a struct so that it's easy for testing
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		},
	).Info("Setting up app info")
	log.Info("Running the server")
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

func main() {
	app := App{
		Name:    "Neverpay",
		Version: "1.0.0",
	}
	err := app.Run()
	if err != nil {
		log.Fatalf("Error starting the server %s", err)
	}
}
