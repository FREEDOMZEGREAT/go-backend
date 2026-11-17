package main

import (
	"errors"
	"fmt"

	"github.com/FREEDOMZEGREAT/go-backend/Day_53/internal/delivery"
	"github.com/FREEDOMZEGREAT/go-backend/Day_53/internal/menu"
	"github.com/FREEDOMZEGREAT/go-backend/Day_53/internal/order"
)

func main() {
	dishes := menu.Menu()

	fmt.Println("=== Order 1: valid, in Bole ===")
	o1 := &order.Order{Area: "Bole"}
	o1.AddItem(dishes[0])
	o1.AddItem(dishes[2])
	total, err := o1.TotalSantim()
	if err != nil {
		fmt.Println("unexpected error:", err)
	} else {
		fmt.Printf("Total: %.2f ETB\n", float64(total)/100)
	}

	fmt.Println("\n=== Order 2: empty ===")
	o2 := &order.Order{Area: "Bole"}
	total, err = o2.TotalSantim()
	if errors.Is(err, order.ErrEmptyOrder) {
		fmt.Println("Rejected: your cart is empty.")
	} else if err != nil {
		fmt.Println("Unexpected error:", err)
	} else {
		fmt.Printf("Total: %.2f ETB\n", float64(total)/100)
	}

	fmt.Println("\n=== Order 3: area Ayat ===")
	o3 := &order.Order{Area: "Ayat"}
	o3.AddItem(dishes[0])
	total, err = o3.TotalSantim()
	if errors.Is(err, delivery.ErrUnservedArea) {
		fmt.Println("Rejected: we don't deliver to Ayat yet.")
	} else if err != nil {
		fmt.Println("Unexpected error:", err)
	} else {
		fmt.Printf("Total: %.2f ETB\n", float64(total)/100)
	}
}
