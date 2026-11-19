package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FREEDOMZEGREAT/go-backend/internal/api"
	"github.com/FREEDOMZEGREAT/go-backend/internal/menu"
)

func main() {
	// 1. Initialize dependencies
	logger := log.New(os.Stdout, "", log.LstdFlags)
	dishStore := menu.NewMemoryStore()

	// 2. Build the application (dependency injection)
	app := api.New(logger, dishStore)

	// 3. Start the server
	logger.Println("Starting Addis Eats API on :8080...")
	err := http.ListenAndServe(":8080", app.Routes())
	if err != nil {
		logger.Fatal(err)
	}
}
