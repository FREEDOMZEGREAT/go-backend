package delivery

import "errors"

// ErrUnservedArea is returned when we don't deliver to the requested area.
var ErrUnservedArea = errors.New("delivery area not served")

// feeByArea maps each servable area to its delivery fee in santim.
var feeByArea = map[string]int{
	"Bole":      3000, // 30.00 ETB
	"Megenagna": 4500, // 45.00 ETB
	"Piassa":    2500, // 25.00 ETB
	"Kazanchis": 2500, // 25.00 ETB
	"Gerji":     3500, // 35.00 ETB
}

// FeeSantim returns the delivery fee for an area, or ErrUnservedArea.
func FeeSantim(area string) (int, error) {
	fee, ok := feeByArea[area]
	if !ok {
		return 0, ErrUnservedArea
	}
	return fee, nil
}