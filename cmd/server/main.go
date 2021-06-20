package main

import (
	"net/http"

	"github.com/1994benc/neverpay-bill-service/internal/database"
	transportHTTP "github.com/1994benc/neverpay-bill-service/internal/transport/http"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func main() {
	app := App{
		Name:    "neverpay-bill-service",
		Version: "1.0.0",
	}
	err := app.Run()
	if err != nil {
		log.Fatalf("Error starting the server %s", err)
		return
	}
}

// Run - runs our application. We set it up in a struct so that it's easy for testing
func (app *App) Run() error {
	app.setUpLogger()
	app.setUpDatabase()

	// TODO: yourService := user.NewService(db)
	// TODO: replace below with handler := transportHTTP.New(yourService)
	handler := transportHTTP.New()
	handler.SetupRoutes()
	err := http.ListenAndServe(":8080", handler.Router)
	return err
}

func (app *App) setUpLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		},
	).Info("Setting up app info")
}

func (app *App) setUpDatabase() (*gorm.DB, error) {
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
	return db, err
}
