package util

import "math"

func RoundFloat(x float64) float64 {
	return math.Round(x*100) / 100
}
