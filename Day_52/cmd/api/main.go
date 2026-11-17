package main

import "fmt"

type Status int

const (
	Placed Status = iota
	Accepted
	Cooking
	OnTheWay
	Delivered
)

func statusName(s Status) string {
	switch s {
	case Placed:
		return "Placed"
	case Accepted:
		return "Accepted"
	case Cooking:
		return "Cooking"
	case OnTheWay:
		return "On the way"
	case Delivered:
		return "Delivered"
	default:
		return "Unknown"
	}
}

func deliveryFee(area string) int {
	switch area {
	case "Bole":
		return 1000
	case "Megenagna":
		return 1500
	case "Kazanchis":
		return 2000
	default:
		return 2000
	}
}

func main() {
	status := Cooking
	dish := "Pasta"
	subtotalSantim := 58500
	serviceRate := 5.0
	service := int(float64(subtotalSantim) * serviceRate / 100)
	delivery := deliveryFee("Bole")
	total := subtotalSantim + service + delivery

	fmt.Println("Current order status:", statusName(status))
	fmt.Println("Dish:", dish)
	fmt.Println("Subtotal:", subtotalSantim, "Santim")
	fmt.Println("Service fee:", service, "Santim")
	fmt.Println("Delivery fee:", delivery, "Santim")
	fmt.Println("Total:", total, "Santim")
}
