package api

import (
	"errors"
	"net/http"

	"github.com/FREEDOMZEGREAT/go-backend/internal/menu"
)

// health returns a simple OK response.
func (app *application) health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// listDishes returns all dishes in the menu.
func (app *application) listDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := app.dishes.All()
	if err != nil {
		app.logger.Print(err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, dishes)
}

// showDish returns a specific dish by ID.
func (app *application) showDish(w http.ResponseWriter, r *http.Request) {
	// Read the {id} wildcard from the URL
	id := r.PathValue("id")

	dish, err := app.dishes.ByID(id)
	switch {
	case errors.Is(err, menu.ErrNotFound):
		writeError(w, http.StatusNotFound, "dish not found")
	case err != nil:
		// Log the unexpected error, return a generic message
		app.logger.Print(err)
		writeError(w, http.StatusInternalServerError, "internal error")
	default:
		writeJSON(w, http.StatusOK, dish)
	}
}
