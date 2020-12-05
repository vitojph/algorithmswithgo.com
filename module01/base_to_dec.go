package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//

func BaseToDec(input string, base int) int {
	const charset = "0123456789ABCDEF"
	var result, power float64
	for i := len(input) - 1; i >= 0; i-- {
		value := float64(strings.Index(charset, string(input[i])))
		result += value * math.Pow(float64(base), power)
		power++
	}
	return int(result)
}
