/*
https://projecteuler.net/problem=9
Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	fmt.Println("problem 9")
	//triplet := 3 + 4 + 5
	triplet := 1000
	fmt.Println("triplet=", triplet)

	//try to break this number into 3 numbers that add up to it
	triplets := list.New()
	for a := 0; a <= triplet; a++ {
		for b := 0; b <= triplet; b++ {
			for c := 0; c <= triplet; c++ {
				if a+b+c == triplet {
					if a*a+b*b == c*c {
						fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
						tmpTriplet := make([]int, 3)
						tmpTriplet[0] = a
						tmpTriplet[1] = b
						tmpTriplet[2] = c
						triplets.PushBack(tmpTriplet)
						fmt.Println("product = ", a*b*c)
					}
				}
			}
		}
	}
	fmt.Println(triplets)
}

func isNaturalRoot(num int) bool {
	sqrt := math.Sqrt(float64(num))
	tmpInt := int(sqrt)
	return sqrt-float64(tmpInt) == 0
}
