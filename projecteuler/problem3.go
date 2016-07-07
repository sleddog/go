//https://projecteuler.net/problem=3
//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("problem 3")
	//create a list of all the prime numbers up to a certain number
	//number := 100
	//number := 13195
	number := 600851475143

	largestPrimeFactor := largestPrimeFactor(number)
	fmt.Printf("Largest Prime Factor of %d, is: %d\n", number, largestPrimeFactor)

	//largestPrime := getLargestPrime(number)
	//fmt.Printf("Largest Prime up to %d, is: %d\n", number, largestPrime)

	/*
		primes := getPrimes(number)
		fmt.Printf("Prime Numbers from 1 - %d\n", number)
		for e := primes.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
		/*

		/*

			primeFactors := primeFactors(number)
			fmt.Printf("Prime Factors for %d:\n", number)
			for e := primeFactors.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
	*/
}

func getLargestPrime(number int) int {
	for i := number / 2; i > 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func getPrimes(number int) *list.List {
	primes := list.New()
	for i := number; i > 2; i-- {
		//fmt.Printf("number %d is ", i)
		if isPrime(i) {
			//fmt.Println("PRIME")
			//fmt.Printf("%d, ", i)
			primes.PushBack(i)
		} else {
			//fmt.Println("NOT PRIME")
		}
	}
	return primes
}

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func largestPrimeFactor(num int) int {
	primes := getPrimes(1000000)
	for e := primes.Front(); e != nil; e = e.Next() {
		intVal, ok := e.Value.(int)
		if !ok {
			fmt.Println("SOMETHINGS NOT RIGHT")
		}
		if num%intVal == 0 {
			fmt.Printf("divisible by %d\n", intVal)
			otherVal := (num / intVal)
			fmt.Printf("%d * %d = %d\n", intVal, otherVal, num)
			//fmt.Printf("---->")
			//more := largestPrimeFactor(otherVal)
			//fmt.Printf("%d <----\n", more)

		}
	}
	/*
		if num%2 == 0 {
			fmt.Println("divisible by 2")
		}
		if num%3 == 0 {
			fmt.Println("divisible by 3")
		}
		if num%5 == 0 {
			fmt.Println("divisible by 5")
		}
		if num%7 == 0 {
			fmt.Println("divisible by 7")
		}
		if num%11 == 0 {
			fmt.Println("divisible by 11")
		}
	*/
	return 0
}

func primeFactors(num int) *list.List {
	primes := getPrimes(num)
	primeFactors := list.New()
	for e := primes.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
		intVal, ok := e.Value.(int)
		if !ok {
			fmt.Println("SOMETHINGS NOT RIGHT")
		}
		if num%intVal == 0 {
			primeFactors.PushBack(e.Value)
		}
	}
	return primeFactors
}
