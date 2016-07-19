/*
https://projecteuler.net/problem=38
Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 38")
	/*
		n := 192
		println(n)
		m := []int{1, 2, 3}
		fmt.Println(m)
		fmt.Println(isPandigitalMultiple(n, m))
		n = 9
		println(n)
		m = []int{1, 2, 3, 4, 5}
		fmt.Println(m)
		fmt.Println(isPandigitalMultiple(n, m))
	*/

	largest := -1
	for n := 1; n < 10000; n++ {
		for i := 2; i < 10; i++ {
			m := []int{}
			for x := 1; x < i; x++ {
				m = append(m, x)
			}
			ok, v := isPandigitalMultiple(n, m)
			if ok {
				fmt.Println("n=", n)
				fmt.Println("m=", m)
				fmt.Println("v=", v)
				if v >= largest {
					largest = v
				}
			}
		}
	}
	fmt.Println("largest=", largest)
}

func isPandigitalMultiple(n int, m []int) (bool, int) {
	s := ""
	for _, v := range m {
		s += strconv.Itoa(n * v)
	}
	if lib.IsPandigital(s) {
		tmpInt, _ := strconv.Atoi(s)
		return true, tmpInt
	}
	return false, 0
}
