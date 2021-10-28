package utils

import "math"

func CheckDivisibleBy(num int32, divisor int32) bool {
	if math.Mod(float64(num), float64(divisor)) == 0 {
		return true
	}
	return false
}
