package main

import (
	"fmt"
	"math"
)

func NegativeOrPositiveCheck(n int, m int) int {
	if n < 0 && m < 0 {
		return 1
	}
	if n < 0 || m < 0 {
		return -1
	}
	return 1
}
func divide(dividend int, divisor int) int {
	sign := NegativeOrPositiveCheck(dividend, divisor)
	dividendPositive := math.Abs(float64(dividend))
	divisorPositive := math.Abs(float64(divisor))
	quotient := 0

	for dividendPositive >= divisorPositive {
		dividendPositive -= divisorPositive
		quotient++
	}
	result := sign * quotient

	return result
}

func main() {
	fmt.Println(divide(10, 3))
}
