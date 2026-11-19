package api

import (
	"log"
	"net/http"

	"github.com/FREEDOMZEGREAT/go-backend/internal/menu"
)

// DishReader is an interface for anything that can provide dishes.
// menu.MemoryStore satisfies this interface implicitly.
type DishReader interface {
	All() ([]menu.Dish, error)
	ByID(id string) (menu.Dish, error)
}

// application holds the shared dependencies for all HTTP handlers.
type application struct {
	logger *log.Logger
	dishes DishReader
}

// New creates a new application instance.
func New(logger *log.Logger, dishes DishReader) *application {
	return &application{
		logger: logger,
		dishes: dishes,
	}
}

// Routes returns the HTTP handler for the application (the ServeMux).
func (app *application) Routes() http.Handler {
	// This will be defined in routes.go, but we declare it here for clarity.
	return app.routes()
}
