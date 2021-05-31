package main

import (
	"log"
)

type App struct{}

// Run - runs our application. We set it up in a struct so that it's easy for testing
func (app *App) Run() error {
	log.Println("Running the server")
	return nil
}

func main() {
	app := App{}
	err := app.Run()
	if err != nil {
		log.Fatalf("Error starting the server %s", err)
	}
}
