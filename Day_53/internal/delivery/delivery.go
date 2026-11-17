package delivery

import "errors"

var ErrUnservedArea = errors.New("unserved area")

var feeByArea = map[string]int{
	"Bole":      1000,
	"Megenagna": 1500,
	"Kazanchis": 2000,
}

func FeeSantim(area string) (int, error) {
	fee, ok := feeByArea[area]
	if !ok {
		return 0, ErrUnservedArea
	}
	return fee, nil
}