/*
https://projecteuler.net/problem=43
Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 43")
	/*
		n := 1406357289
		fmt.Println(n)
		fmt.Println(hasProperty(n))
	*/
	sum := 0
	for _, n := range lib.Permutations("0123456789") {
		if hasProperty(n) {
			fmt.Println(n)
			i, _ := strconv.Atoi(n)
			sum += i

		}

	}
	fmt.Println("sum=", sum)
}

func hasProperty(s string) bool {
	//first check its pandigital
	if !lib.IsPandigital10(s) {
		return false
	}
	p := []int{2, 3, 5, 7, 11, 13, 17}
	for i := 0; i < len(p); i++ {
		subStr := s[i+1 : i+4]
		tmpInt, _ := strconv.Atoi(subStr)
		if tmpInt%p[i] != 0 {
			return false
		}
	}
	return true
}
