/*
https://projecteuler.net/problem=16
Power digit sum

215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("problem 16")
	value := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	str := value.String()
	fmt.Println(str)
	sum := 0
	for i := 0; i < len(str); i++ {
		tmpInt, _ := strconv.Atoi(str[i : i+1])
		sum += tmpInt
	}
	fmt.Println("sum=", sum)
}
