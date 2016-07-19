/*
https://projecteuler.net/problem=39
Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("problem 39")
	max := -1
	maxP := -1
	for p := 1; p <= 1000; p++ {
		solutions := [][]int{}
		for b := 1; b < p/2; b++ {
			for a := 1; a < b; a++ {
				fc := pythag(a, b)
				if isWholeNum(fc) {
					c := int(fc)
					if a+b+c == p {
						//fmt.Printf("a:%v, b:%v, c:%v\n", a, b, c)
						solutions = append(solutions, []int{a, b, c})
					}
				}
			}
		}
		l := len(solutions)
		if l > max {
			max = l
			maxP = p
			fmt.Printf("p=%v: %v\n", p, solutions)
		}
	}
	fmt.Println("max=", max)
	fmt.Println("maxP=", maxP)
}

func pythag(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func isWholeNum(a float64) bool {
	return a == float64(int64(a))
}
