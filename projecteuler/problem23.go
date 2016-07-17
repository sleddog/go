/*
https://projecteuler.net/problem=23
Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

package main

import (
	"./lib"
	"fmt"
)

var abundantNums []int

func main() {
	fmt.Println("problem 23")
	max := 28123
	abundantNums = abundantNumbers(max)
	sum := 0
	for i := 1; i <= max; i++ {
		if cannotBeWritten(i) {
			sum += i
		}
	}
	fmt.Println("sum=", sum)
}

//check if this number can be formed from 2 of the abudant numbers
func cannotBeWritten(i int) bool {
	for _, a := range abundantNums {
		if a < i {
			for _, a2 := range abundantNums {
				if a2 <= i-a {
					if a+a2 == i {
						return false
					}

				} else {
					break
				}
			}
		} else {
			break
		}
	}
	return true
}

func abundantNumbers(max int) []int {
	//find all abundant numbers up to max
	nums := []int{}
	for i := 1; i <= max; i++ {
		if isAbundant(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func isAbundant(n int) bool {
	divisors := lib.ProperDivisors(n)
	sum := 0
	for _, v := range divisors {
		sum += v
	}
	if sum > n {
		return true
	}
	return false
}
