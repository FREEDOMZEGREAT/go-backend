package order

import (
	"errors"
	"fmt"

	"github.com/FREEDOMZEGREAT/go-backend/internal/delivery"
	"github.com/FREEDOMZEGREAT/go-backend/internal/menu"
)

// Sentinel errors for the order package
var (
	ErrEmptyOrder = errors.New("order has no items")
)

// Order represents a customer's order.
type Order struct {
	Items []menu.Dish
	Area  string
}

// AddItem appends a dish to the order.
// Pointer receiver because it modifies the receiver.
func (o *Order) AddItem(d menu.Dish) {
	o.Items = append(o.Items, d)
}

// SubtotalSantim returns the sum of item prices, before fees.
// Value receiver because it only reads.
func (o Order) SubtotalSantim() int {
	subtotal := 0
	for _, d := range o.Items {
		subtotal += d.PriceSantim
	}
	return subtotal
}

// TotalSantim returns the total including service fee and delivery.
func (o Order) TotalSantim() (int, error) {
	if len(o.Items) == 0 {
		return 0, ErrEmptyOrder
	}

	subtotal := o.SubtotalSantim()

	// Service fee: 5% of subtotal, truncated to santim
	serviceFee := int(float64(subtotal) * 5.0 / 100)

	// Free delivery if subtotal >= 50000 santim (500 ETB)
	const freeDeliveryThreshold = 50000
	deliveryFee := 0
	if subtotal < freeDeliveryThreshold {
		fee, err := delivery.FeeSantim(o.Area)
		if err != nil {
			// Wrap the delivery error so callers can use errors.Is
			return 0, fmt.Errorf("compute delivery fee: %w", err)
		}
		deliveryFee = fee
	}

	return subtotal + serviceFee + deliveryFee, nil
}
