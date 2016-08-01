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

var p []int

func main() {
	fmt.Println("problem 47")
	size := 651
	p = lib.GetPrimesSieve(size)

	//s := distinctTwoPrimeFactors(size)
	//fmt.Println(s)

	s := distinctThreePrimeFactors(size)
	fmt.Println(s)
}

func distinctTwoPrimeFactors(size int) string {
	d := 2
	fmt.Printf("looking for first %v consecutive numbers to have %v distinct prime factors\n", d, d)

	lastFound := 0
	s := ""

	for i := 1; i < size; i++ {
		for _, v1 := range p {
			for _, v2 := range p {
				if v1 >= v2 {
					continue
				}
				if v1*v2 == i {
					if lastFound+1 == i {
						s += fmt.Sprintf("%v * %v = %v\n", v1, v2, i)
						return s
					}
					s = fmt.Sprintf("%v * %v = %v\n", v1, v2, i)
					lastFound = i
				}
			}
		}
	}
	return "?"
}

func distinctThreePrimeFactors(size int) string {
	d := 3
	fmt.Printf("looking for first %v consecutive numbers to have %v distinct prime factors\n", d, d)

	count := 0
	prev := -1
	s := ""
	product := 0

	for i := 1; i < size; i++ {
		for _, v1 := range p {
			for _, v2 := range p {
				if v1 >= v2 {
					continue
				}
				for _, v3 := range p {
					if v2 >= v3 {
						continue
					}
					//to capture 644... need to square each term as well???
					product = v1 * v2 * v3
					if product == i {
						fmt.Printf("i==%v   -    %v * %v * %v = %v\n", i, v1, v2, v3, product)
						if prev == i-1 {
							count++
							s += fmt.Sprintf("%v * %v * %v = %v\n", v1, v2, v3, i)
						} else {
							count = 1
							s = fmt.Sprintf("%v * %v * %v = %v\n", v1, v2, v3, i)
						}
						prev = i

						if count == d {
							return s
						}
					}
				}
			}
		}
	}
	return "?"
}
