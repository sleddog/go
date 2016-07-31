/*
https://projecteuler.net/problem=46
<NAME>

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2*1^2
15 = 7 + 2*2^2
21 = 3 + 2*3^2
25 = 7 + 2*3^2
27 = 19 + 2*2^2
33 = 31 + 2*1^2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
*/

package main

import (
	"./lib"
	"fmt"
	"math"
)

var p []int

func main() {
	fmt.Println("problem 46")
	size := 100000
	p = lib.GetPrimesSieve(size)
	for i := 9; i < size; i++ {
		if isOddComposite(i) {
			b, s := isSumPrimeTwiceSquare(i)
			if b {
				fmt.Printf("%v = %v\n", i, s)
			} else {
				fmt.Printf("FOUND IT i=%v, %s\n", i, s)
				break
			}
		}

	}
}

func isOddComposite(n int) bool {
	if n%2 == 0 {
		return false
	}
	if lib.IsPrime(n) {
		return false
	}
	return true
}

func isSumPrimeTwiceSquare(n int) (bool, string) {
	for _, v := range p {
		for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
			sum := v + 2*(i*i)
			if sum == n {
				s := fmt.Sprintf("%v + 2 * %v^2", v, i)
				return true, s
			}
		}
	}
	return false, "can't find it...  goldbach was wrong!"
}
