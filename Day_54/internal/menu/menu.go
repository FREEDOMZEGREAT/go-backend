package menu

import "errors"

// Sentinel errors for the menu package
var (
	ErrNotFound = errors.New("dish not found")
)

// Dish represents a single item on the menu.
type Dish struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PriceSantim int    `json:"price_santim"`
	IsFasting   bool   `json:"is_fasting"`
}

// seedDishes returns the initial set of dishes used by the memory store.
// It is unexported — only code inside this package should use it.
func seedDishes() []Dish {
	return []Dish{
		{ID: "d1", Name: "Doro Wat", PriceSantim: 45000, IsFasting: false},
		{ID: "d2", Name: "Tibs", PriceSantim: 52000, IsFasting: false},
		{ID: "d3", Name: "Shiro", PriceSantim: 18000, IsFasting: true},
		{ID: "d4", Name: "Misir Wat", PriceSantim: 20000, IsFasting: true},
		{ID: "d5", Name: "Kitfo", PriceSantim: 48000, IsFasting: false},
		{ID: "d6", Name: "Fasolia", PriceSantim: 22000, IsFasting: true},
		{ID: "d7", Name: "Beyaynetu", PriceSantim: 25000, IsFasting: true},
		{ID: "d8", Name: "Gored Gored", PriceSantim: 55000, IsFasting: false},
	}
}