/*
https://projecteuler.net/problem=10
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"./lib"
	"fmt"
)

func main() {
	fmt.Println("problem 10")
	primes := lib.GetPrimes(2000000)
	sum := 0
	for e := primes.Front(); e != nil; e = e.Next() {
		intVal, ok := e.Value.(int)
		if !ok {
			fmt.Println("SOMETHINGS NOT RIGHT")
		}
		fmt.Println(intVal)
		sum = sum + intVal
	}
	fmt.Println("sum=", sum)
}
