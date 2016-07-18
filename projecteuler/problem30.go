/*
https://projecteuler.net/problem=30
Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("problem 30")
	//nums := digitPowers(4)
	nums := digitPowers(5)
	fmt.Println(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum=", sum)
}

func digitPowers(n int) []int {
	nums := []int{}
	for i := 2; i < 10000000; i++ {
		s := strconv.Itoa(i)
		//fmt.Println(s)
		sum := 0.0
		for j := 0; j < len(s); j++ {
			x, _ := strconv.Atoi(s[j : j+1])
			//fmt.Println("x=", x)
			v := math.Pow(float64(x), float64(n))
			//fmt.Printf("%v^%v=%v\n", x, n, v)
			sum += v
		}
		if sum == float64(i) {
			//fmt.Println("FOUND ONE", i)
			nums = append(nums, i)
		}
	}
	return nums
}
