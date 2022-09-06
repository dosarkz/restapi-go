package floats

import (
	"fmt"
	"math"
	"strconv"
)

func ToStandardFloat(val float64, precision uint) (*float64, error) {
	ratio := math.Pow(10, float64(precision))
	roundedFloat := math.Round(val*ratio) / ratio
	s := fmt.Sprintf("%.3f", roundedFloat)
	if a, err := strconv.ParseFloat(s, 64); err == nil {
		return &a, nil
	} else {
		return nil, err
	}
}
