/*
https://projecteuler.net/problem=57
Square root convergents

It is possible to show that the square root of two can be expressed as an infinite continued fraction.

√ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...

By expanding this for the first four iterations, we get:

1 + 1/2 = 3/2 = 1.5
1 + 1/(2 + 1/2) = 7/5 = 1.4
1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408, but the eighth expansion, 1393/985, is the first example where the number of digits in the numerator exceeds the number of digits in the denominator.

In the first one-thousand expansions, how many fractions contain a numerator with more digits than denominator?
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("problem 57")
	/*
		//first
		n, d := expand(1, 1, 2)
		fmt.Println(n, d)

		//second
		n, d = expand(2, 1, 2)
		n, d = expand(1, d, n)
		fmt.Println(n, d)

		//third
		n, d = expand(2, 1, 2)
		n, d = expand(2, d, n)
		n, d = expand(1, d, n)
		fmt.Println(n, d)

		//fourth
		d = 1
		n = 2
		for i := 0; i < 3; i++ {
			n, d = expand(2, d, n)
		}
		n, d = expand(1, d, n)
		fmt.Println(n, d)

		//fifth
		fmt.Println(testExpand(5))
	*/

	moreDigitsCount := 0
	for i := 1; i <= 1000; i++ {
		if testExpand(i) {
			fmt.Println("FOUND ONE: ", i)
			moreDigitsCount++
		}
	}
	fmt.Println("moreDigitsCount=", moreDigitsCount)
}

func testExpand(times int) bool {
	d := big.NewInt(int64(1))
	n := big.NewInt(int64(2))
	for i := 0; i < times-1; i++ {
		n, d = expand(big.NewInt(int64(2)), d, n)
	}
	n, d = expand(big.NewInt(int64(1)), d, n)
	fmt.Println(n, d)
	fmt.Println("~=", new(big.Float).Quo(new(big.Float).SetInt(n), new(big.Float).SetInt(d)))

	//check length of digits
	nStr := n.String()
	dStr := d.String()
	return len(nStr) > len(dStr)
}

func expand(x, n, d *big.Int) (*big.Int, *big.Int) {
	//return x*d + n, d
	tmp := new(big.Int).Mul(x, d)
	return new(big.Int).Add(tmp, n), d
}
