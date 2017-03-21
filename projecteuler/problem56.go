/*
https://projecteuler.net/problem=56
Powerful digit sum

A googol (10^100) is a massive number: one followed by one-hundred zeros; 100^100 is almost unimaginably large: one followed by two-hundred zeros. Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, a^b, where a, b < 100, what is the maximum digital sum?
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("problem 56")
	max := 0
	maxA := -1
	maxB := -1
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			value := new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			//fmt.Printf("a=%d, b=%d\n", a, b)
			//fmt.Println("a^b=", value)
			i := sumDigits(value)
			if i > max {
				max = i
				maxA = a
				maxB = b
			}
		}
	}
	fmt.Println("max=", max)
	fmt.Println("a=", maxA)
	fmt.Println("b=", maxB)
}

func sumDigits(n *big.Int) int {
	str := n.String()
	sum := 0
	for _, b := range str {
		tmpInt, _ := strconv.Atoi(string(b))
		sum += tmpInt
	}
	return sum
}
