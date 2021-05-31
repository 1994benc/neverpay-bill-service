package main

import (
	transportHTTP "1994benc/neverpay-api/internal/transport/http"
	"log"
	"net/http"
)

type App struct{}

// Run - runs our application. We set it up in a struct so that it's easy for testing
func (app *App) Run() error {
	log.Println("Running the server")
	handler := transportHTTP.New()
	handler.SetupRoutes()
	err := http.ListenAndServe(":8080", handler.Router)
	return err
}

func main() {
	app := App{}
	err := app.Run()
	if err != nil {
		log.Fatalf("Error starting the server %s", err)
	}
}
