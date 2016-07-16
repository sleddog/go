package lib

import (
	"math"
)

func Factors(num int) []int {
	factors := []int{}
	for k := 1; k <= int(math.Sqrt(float64(num))); k++ {
		if num%k == 0 {
			factors = append(factors, k)
			divisor := num / k
			//only add the divisor if it isn't the same
			if divisor != k {
				factors = append(factors, divisor)
			}
		}
	}
	return factors
}

func Factors2(num int) []int {
	factors := []int{}
	for k := 1; k <= num; k++ {
		if num%k == 0 {
			factors = append(factors, k)
		}
	}
	return factors
}
