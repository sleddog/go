//https://projecteuler.net/problem=5
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 5")
	fmt.Println(smallestNumber(1, 10))
	number := 2520

	fmt.Println("check number", number)
	fmt.Println(checkNumber(number, 1, 10))

	for i := 1; i < 100000000000; i++ {
		//fmt.Println("i=", i)
		if checkNumber(i, 1, 20) {
			fmt.Println(i)
			break
		}
		/*
			if i%100000 == 0 {
				fmt.Print('.', i)
			}
		*/
	}
}

func smallestNumber(start, end int) int {
	return 0
}

func checkNumber(number, start, end int) bool {
	for i := start; i < end; i++ {
		if number%i != 0 {
			return false
		}
	}
	return true
}
