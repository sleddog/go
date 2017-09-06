/*
https://projecteuler.net/problem=63
Powerful digit counts

The 5-digit number, 16807=7^5, is also a fifth power. Similarly, the 9-digit number, 134217728=8^9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("problem 63")
	fmt.Println(exp(7, 5))
	fmt.Println(exp(8, 9))

	count := 0
	for x := 1; x < 50; x++ {
		for y := 1; y < 1000; y++ {
			z := exp(x, y)
			zStr := z.String()
			length := len(zStr)
			if length == y {
				fmt.Println("FOUND ONE", x, "^", y, "=", zStr)
				count = count + 1
			}
		}
	}
	fmt.Println("count = ", count)
}

func exp(x, y int) *big.Int {
	value := new(big.Int).Exp(big.NewInt(int64(x)), big.NewInt(int64(y)), nil)
	return value
}
