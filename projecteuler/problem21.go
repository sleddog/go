/*
https://projecteuler.net/problem=21
Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"./lib"
	"fmt"
)

func main() {
	fmt.Println("problem 21")
	sum := 0
	for a := 1; a <= 10000; a++ {
		b := d(a)
		a2 := d(b)
		if a == a2 && a != b {
			fmt.Printf("d(%d) = d(%d)\n", a, b)
			sum += a
			//sum += b  this was double counting...
		}
	}
	fmt.Println("sum=", sum)
}

func d(n int) int {
	divisors := lib.ProperDivisors(n)
	sum := 0
	for _, val := range divisors {
		sum += val
	}
	return sum
}
