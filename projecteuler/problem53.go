/*
https://projecteuler.net/problem=53
Combinatoric selections

There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, C(5,3) = 10.

In general:

C(n,r) = n! / (r!(n−r)!)

where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.

It is not until n = 23, that a value exceeds one-million: C(23,10) = 1144066.

How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are greater than one-million?
*/

package main

import (
	"./lib"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("problem 53")
	//n := 5
	//r := 3
	//n := 23
	//r := 10
	//C(n, r)

	million := big.NewInt(1000000)

	count := 0
	for n := 1; n <= 100; n++ {
		for r := 0; r < n; r++ {
			i := C(n, r)
			compare := i.Cmp(million)
			if compare >= 0 {
				fmt.Printf("C(%d,%d)=", n, r)
				fmt.Println(i)
				count++
			}
		}
	}
	fmt.Println("count=", count)
}

/*
In general:

C(n,r) = n! / (r!(n−r)!)

where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.
*/
func C(n, r int) *big.Int {
	if r > n {
		fmt.Println("error... r <= n!")
		return big.NewInt(-1)
	}
	t := lib.FactorialBigInt(n)
	b := new(big.Int)
	b.Mul(lib.FactorialBigInt(r), lib.FactorialBigInt(n-r))
	result := new(big.Int)
	result.Div(t, b)
	return result
}
