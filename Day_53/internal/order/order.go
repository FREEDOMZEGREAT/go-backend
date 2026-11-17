package order

import (
	"errors"
	"fmt"

	"github.com/FREEDOMZEGREAT/go-backend/Day_53/internal/delivery"
	"github.com/FREEDOMZEGREAT/go-backend/Day_53/internal/menu"
)

var ErrEmptyOrder = errors.New("empty order")

type Order struct {
	Dishes []menu.Dish
	Area   string
}

func (o *Order) AddItem(d menu.Dish) {
	o.Dishes = append(o.Dishes, d)
}

func (o *Order) SubtotalSantim() int {
	subtotal := 0
	for _, d := range o.Dishes {
		subtotal += d.PriceSantim
	}
	return subtotal
}

func (o *Order) TotalSantim() (int, error) {
	if len(o.Dishes) == 0 {
		return 0, ErrEmptyOrder
	}
	subtotal := o.SubtotalSantim()
	serviceFee := int(float64(subtotal) * 0.05)
	const freeDeliveryThreshold = 5000
	deliveryFee := 0
	if subtotal < freeDeliveryThreshold {
		fee, err := delivery.FeeSantim(o.Area)
		if err != nil {
			return 0, fmt.Errorf("compute delivery fee: %w", err)
		}
		deliveryFee = fee
	}
	return subtotal + serviceFee + deliveryFee, nil
}