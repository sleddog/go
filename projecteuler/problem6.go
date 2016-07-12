/*
https://projecteuler.net/problem=6



The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 6")
	number := 100
	sum := sumOfSquares(number)
	square := squareOfSum(number)
	diff := square - sum
	fmt.Printf("sum of squares for %d: %d\n", number, sum)
	fmt.Printf("sum of squares for %d: %d\n", number, square)
	fmt.Printf("diff: %d\n", diff)
}

func sumOfSquares(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i
	}
	return sum * sum
}
