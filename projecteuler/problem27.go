/*
https://projecteuler.net/problem=27
Quadratic primes

Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |−4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/

package main

import (
	"./lib"
	"fmt"
)

func main() {
	fmt.Println("problem 27")
	/*
		fmt.Println("euler's:")
		for n := 0; n < 40; n++ {
			v := eulers(n)
			b := lib.IsPrime(v)
			fmt.Printf("n=%d: %d - %t\n", n, v, b)
		}
		fmt.Println("incredible:")
		for n := 0; n < 80; n++ {
			v := incredible(n)
			b := lib.IsPrime(v)
			fmt.Printf("n=%d: %d - %t\n", n, v, b)
		}
	*/

	//lets find it...
	maxConsecutive := -1
	bestA := -1
	bestB := -1
	for a := -1000; a < 1001; a++ {
		for b := -1000; b < 1001; b++ {
			consecutiveCount := 0
			n := 0
			for {
				v := tester(n, a, b)
				t := lib.IsPrime(v)
				if t {
					consecutiveCount++
					n++
				} else {
					if consecutiveCount >= maxConsecutive {
						fmt.Printf("n=%d, a=%d, b=%d: %d - %t\n", n, a, b, v, t)
						maxConsecutive = consecutiveCount
						bestA = a
						bestB = b
					}
					break
				}
			}
		}
	}
	fmt.Println("maxConsecutive=", maxConsecutive)
	fmt.Println("bestA=", bestA)
	fmt.Println("bestB=", bestB)
	fmt.Println("product=", bestA*bestB)
}

func eulers(n int) int {
	return n*n + n + 41
}

func incredible(n int) int {
	return n*n - 79*n + 1601
}

func tester(n, a, b int) int {
	return (n * n) + (a * n) + b
}
