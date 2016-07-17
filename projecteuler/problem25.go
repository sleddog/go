/*
https://projecteuler.net/problem=25
1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("problem 25")
	i := 1
	for {
		s := fib(big.NewInt(int64(i))).String()
		l := len(s)
		if l >= 1000 {
			fmt.Printf("%s - %d\n", s, l)
			fmt.Println("index = ", i)
			break
		}
		i++
	}
}

func fib(n *big.Int) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(n) < 0; i.Add(i, big.NewInt(1)) {
		a, b = b, a.Add(a, b)
	}
	return a
}
