/*
https://projecteuler.net/problem=40
Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...
             ^

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("problem 40")
	s := ""
	e := 0
	i := 1
	n := 1 //10^e==10^0
	product := 1
	for {
		s += strconv.Itoa(i)
		l := len(s)
		if l >= n {
			dn := s[n-1 : n]
			fmt.Printf("d%d = %v\n", n, dn)
			tmpInt, _ := strconv.Atoi(dn)
			product *= tmpInt

			e++
			if e > 6 {
				//we done
				break
			}
			n = int(math.Pow(float64(10), float64(e)))
		}

		i++
	}
	fmt.Println("product = ", product)
}
