package lib

import (
	"math"
	//	"sort"
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

func ProperDivisors(n int) []int {
	//basically just factors minus the term 'n'
	factors := Factors(n)
	proper := make([]int, 0)
	for i := 0; i < len(factors); i++ {
		if factors[i] != n {
			proper = append(proper, factors[i])
		}

	}
	//might as well sort it too... (probably a performance hit)
	//tested it... and its a big hit...
	//sort.Ints(proper)
	return proper
}

func ProperDivisors2(n int) []int {
	d := make([]int, 0)
	for i := 1; i < n; i++ {
		if n%i == 0 {
			d = append(d, i)
		}
	}
	return d
}

//needed for problem47
func DistinctPrimeFactors(i int) []int {
	pf := []int{}
	factors := make(map[int]int)
	for n := i; n > 1; {
		factor := getFactor(n)
		if f, ok := factors[factor]; ok {
			factors[factor] = f + 1
		} else {
			factors[factor] = 1
		}
		n = n / factor
	}
	//convert map to a slice, by raising the key to the value (k^v)
	for k, v := range factors {
		pf = append(pf, int(math.Pow(float64(k), float64(v))))
	}
	return pf
}

func getFactor(n int) int {
	x := 2
	for x < int(math.Sqrt(float64(n)))+1 {
		if n%x == 0 {
			return x
		}
		x += 1
	}
	return n
}
