package api

import (
	"net/http"
)

// routes sets up the ServeMux and registers all handlers.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("GET /health", app.health)

	// Menu endpoints
	mux.HandleFunc("GET /menu", app.listDishes)
	mux.HandleFunc("GET /menu/{id}", app.showDish)

	// Note: Order endpoints (POST /orders) will be added on Day 55.

	return mux
}