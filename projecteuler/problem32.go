/*
https://projecteuler.net/problem=32
Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 32")
	//39 × 186 = 7254
	//fmt.Println(isPandigital(39, 186, 7254))
	m := make(map[int]bool)
	count := 0
	for i := 1; i < 10000; i++ {
		for j := 1; j < 10000; j++ {
			p := i * j
			//fmt.Printf("%v * %v = %v\n", i, j, p)
			if isPandigital(i, j, p) {
				fmt.Println("FOUND ONE", i, j, p)
				count++
				m[p] = true
			}
		}
	}
	fmt.Println(count)
	fmt.Println(len(m))
	fmt.Println(m)
	sum := 0
	for key, _ := range m {
		fmt.Println(key)
		sum += key
	}
	fmt.Println("sum=", sum)
}

func isPandigital(i, j, p int) bool {
	x := strconv.Itoa(i)
	y := strconv.Itoa(j)
	z := strconv.Itoa(p)
	totalLen := len(x) + len(y) + len(z)

	if totalLen != 9 {
		return false
	}

	m := make(map[int]bool)
	//now check every number
	for i := 0; i < len(x); i++ {
		tmpInt, _ := strconv.Atoi(x[i : i+1])
		if tmpInt == 0 {
			return false
		}
		if m[tmpInt] {
			return false
		}
		m[tmpInt] = true
	}
	for i := 0; i < len(y); i++ {
		tmpInt, _ := strconv.Atoi(y[i : i+1])
		if tmpInt == 0 {
			return false
		}
		if m[tmpInt] {
			return false
		}
		m[tmpInt] = true
	}
	for i := 0; i < len(z); i++ {
		tmpInt, _ := strconv.Atoi(z[i : i+1])
		if tmpInt == 0 {
			return false
		}
		if m[tmpInt] {
			return false
		}
		m[tmpInt] = true
	}
	if len(m) == 9 {
		return true
	} else {
		return false
	}
}
