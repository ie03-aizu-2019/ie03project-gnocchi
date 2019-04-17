package utils

import "math"

func Round(f float64, n int) float64 {
	return math.Floor(f*math.Pow10(n-1)+0.5) / math.Pow10(n-1)
}

func NearEqual(a float64, b float64) bool {
	eps := 10e-7
	t := a - b

	return -eps <= t && t <= eps
}
