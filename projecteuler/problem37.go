/*
https://projecteuler.net/problem=37
Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 37")
	//p := 3797
	//print(isTruncatablePrime(p))
	primes := lib.GetPrimesSieve(1000000)
	sum := 0
	count := 0
	for _, v := range primes {
		if isTruncatablePrime(v) {
			println(v)
			sum += v
			count++
		}
	}
	fmt.Println("count=", count)
	fmt.Println("sum=", sum)
}

func isTruncatablePrime(n int) bool {
	//2, 3, 5, and 7 are not considered to be truncatable primes.
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return false
	}
	if !lib.IsPrime(n) {
		return false
	}
	s := strconv.Itoa(n)
	l := len(s)
	for i := 0; i < l; i++ {
		right, _ := strconv.Atoi(s[i:l])
		if !lib.IsPrime(right) {
			return false
		}
		left, _ := strconv.Atoi(s[0 : l-i])
		if !lib.IsPrime(left) {
			return false
		}
	}
	return true
}
