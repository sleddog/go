/*
https://projecteuler.net/problem=20
Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("problem 20")
	intVal := factorial(100)
	strVal := intVal.String()
	fmt.Println(strVal)
	fmt.Println("sum=", sumDigits(strVal))
}

func sumDigits(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		tmpInt, _ := strconv.Atoi(s[i : i+1])
		sum += tmpInt
	}
	return sum
}

func factorial(n int) *big.Int {
	product := big.NewInt(1)
	for i := 1; i <= n; i++ {
		product = product.Mul(product, big.NewInt(int64(i)))
	}
	return product
}
