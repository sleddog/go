/*
https://projecteuler.net/problem=33
Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 33")
	top := 1
	bottom := 1
	for i := 10; i < 100; i++ {
		for j := 10; j < 100; j++ {
			if i < j {
				//fraction
				f := float64(i) / float64(j)
				//fmt.Printf("%v / %v = %v\n", i, j, f)

				//remove the same digit
				si := strconv.Itoa(i)
				i1, _ := strconv.Atoi(si[0:1])
				i2, _ := strconv.Atoi(si[1:2])
				//fmt.Printf("i = %v %v\n", i1, i2)
				sj := strconv.Itoa(j)
				j1, _ := strconv.Atoi(sj[0:1])
				j2, _ := strconv.Atoi(sj[1:2])
				//fmt.Printf("j = %v %v\n", j1, j2)

				if i2 == 0 && j2 == 0 {
					//trivial example
					continue
				}

				//try the 4 ways
				if i1 == j1 {
					newF := float64(i2) / float64(j2)
					if newF == f {
						fmt.Println("FOUND ONE")
						fmt.Printf("%v / %v = %v\n", i2, j2, newF)
						fmt.Printf("%v / %v = %v\n", i, j, f)
						top *= i2
						bottom *= j2
					}
				}
				if i1 == j2 {
					newF := float64(i2) / float64(j1)
					if newF == f {
						fmt.Println("FOUND ONE")
						fmt.Printf("%v / %v = %v\n", i2, j1, newF)
						fmt.Printf("%v / %v = %v\n", i, j, f)
						top *= i2
						bottom *= j1
					}
				}
				if i2 == j1 {
					newF := float64(i1) / float64(j2)
					if newF == f {
						fmt.Println("FOUND ONE")
						fmt.Printf("%v / %v = %v\n", i1, j2, newF)
						fmt.Printf("%v / %v = %v\n", i, j, f)
						top *= i1
						bottom *= j2
					}
				}
				if i2 == j2 {
					newF := float64(i1) / float64(j1)
					if newF == f {
						fmt.Println("FOUND ONE")
						fmt.Printf("%v / %v = %v\n", i1, j1, newF)
						fmt.Printf("%v / %v = %v\n", i, j, f)
						top *= i1
						bottom *= j1
					}
				}
			}
		}
	}
	fmt.Println("top =", top)
	fmt.Println("bottom =", bottom)
	fmt.Println("reduce to common demoninator ^^")
}
