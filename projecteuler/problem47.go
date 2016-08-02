/*
https://projecteuler.net/problem=47
Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 * 7
15 = 3 * 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2^2 * 7 * 23
645 = 3 * 5 * 43
646 = 2 * 17 * 19

Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
*/

package main

import (
	"./lib"
	"fmt"
)

func main() {
	fmt.Println("problem 47")
	size := 1000000
	consec := 4
	prev := -1
	inARow := 0

	for n := 1; n < size; n++ {
		pf := lib.DistinctPrimeFactors(n)
		if len(pf) == consec {
			fmt.Println(n, pf)
			if prev+1 == n {
				inARow++
			} else {
				inARow = 1
			}
			prev = n
		}
		if inARow == consec {
			fmt.Println("GOT IT")
			return
		}
	}
}
