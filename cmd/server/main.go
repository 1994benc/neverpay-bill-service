package main

import (
	"net/http"

	"github.com/1994benc/minimal-go-microservice-template/internal/database"
	transportHTTP "github.com/1994benc/minimal-go-microservice-template/internal/transport/http"

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
	// TODO: userService := user.NewService(db)
	handler := transportHTTP.New() // TODO: pass in userService like handler := transportHTTP.New(userService)
	handler.SetupRoutes()
	err = http.ListenAndServe(":8080", handler.Router)
	return err
}

func main() {
	app := App{
		Name:    "minimal-go-microservice-template",
		Version: "1.0.0",
	}
	err := app.Run()
	if err != nil {
		log.Fatalf("😢 Error starting the server %s", err)
	} else {
		log.Println("***** 😀 Sucessfully started the server!!! 🙌 *****")
	}
}
