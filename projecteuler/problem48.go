/*
https://projecteuler.net/problem=48
Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("problem 48")
	sum := new(big.Int)
	for i := 1; i <= 1000; i++ {
		value := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil)
		sum.Add(sum, value)
	}
	bigStr := sum.String()
	fmt.Println(bigStr)
	fmt.Println("last 10 digits of bigStr:")
	fmt.Println(bigStr[len(bigStr)-10 : len(bigStr)])
}
