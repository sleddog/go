/*
https://projecteuler.net/problem=34
Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 34")
	//i := 145
	//fmt.Println(isCurious(i))
	sum := 0
	for i := 3; i < 1000000000; i++ {
		if isCurious(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("sum=", sum)
}

func isCurious(i int) bool {
	s := strconv.Itoa(i)
	sum := 0
	for j := 0; j < len(s); j++ {
		tmpInt, _ := strconv.Atoi(s[j : j+1])
		tmpF := factorial(tmpInt)
		sum += tmpF
	}

	return i == sum
}

func factorial(n int) int {
	product := 1
	for i := 1; i <= n; i++ {
		product = product * i
	}
	return product
}
