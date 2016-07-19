/*
https://projecteuler.net/problem=41
Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 41")
	for v := 1000000000; v > 3; v-- {
		if v%100000 == 0 {
			fmt.Printf("%v\n", v)
		}
		n := len(strconv.Itoa(v))
		if lib.IsPrime(v) {
			//fmt.Printf("prime %v\n", v)
			if lib.IsNPandigitalInt(n, v) {
				fmt.Printf("%v : %v\n", n, v)
				break
			}
		}
	}

}
