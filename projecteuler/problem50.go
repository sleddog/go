/*
https://projecteuler.net/problem=50
Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

package main

import (
	"./lib"
	"fmt"
)

func main() {
	fmt.Println("problem 50")
	p := lib.GetPrimesFromCache(1000000)
	mostPrimes := 0
	bestPrime := 0
	for i := 0; i < len(p); i++ {
		sum := 0
		consec := []int{}
		for j := i; j < len(p); j++ {
			sum += p[j]
			consec = append(consec, p[j])
			if len(consec) > mostPrimes {
				if lib.IsPrime(sum) {
					//fmt.Println(sum, len(consec))
					mostPrimes = len(consec)
					bestPrime = sum
				}
			}
			if sum > 1000000 {
				break
			}
		}
	}
	fmt.Printf("bestPrime=%v, len=%v\n", bestPrime, mostPrimes)
}
