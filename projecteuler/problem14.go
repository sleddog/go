/*
https://projecteuler.net/problem=14
Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 14")
	longestChain := 0
	startNumber := 1
	for n := 1; n < 1000000; n++ {
		chain := CollatzChain(n)
		if len(chain) >= longestChain {
			longestChain = len(chain)
			startNumber = n
		}
	}
	fmt.Println(longestChain)
	fmt.Println(startNumber)
	fmt.Println(CollatzChain(startNumber))
}

func CollatzChain(n int) []int {
	chain := []int{}
	for {
		chain = append(chain, n)
		if n <= 1 {
			break
		}
		n = Collatz(n)
	}
	return chain
}

func Collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
