/*
https://projecteuler.net/problem=47
Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 * 7
15 = 3 * 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2 * 7 * 23
645 = 3 * 5 * 43
646 = 2 * 17 * 19

Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
*/

package main

import (
	"./lib"
	"fmt"
)

var p []int

func main() {
	fmt.Println("problem 47")
	size := 16
	p = lib.GetPrimesSieve(size)
	fmt.Println(p)
	distinctTwoPrimeFactors(size)
}

func distinctTwoPrimeFactors(size int) {
	d := 2
	fmt.Printf("looking for first %v consecutive numbers to have %v distinct prime factors\n", d, d)

	lastFound := 0

	for i := 1; i < size; i++ {
		for _, v1 := range p {
			for _, v2 := range p {
				if v1 >= v2 {
					continue
				}
				if v1*v2 == i {
					fmt.Printf("%v * %v = %v\n", v1, v2, i)
					if lastFound+1 == i {
						fmt.Println("consecutive")
						return
					}
					lastFound = i
				}
			}
		}
	}
}
