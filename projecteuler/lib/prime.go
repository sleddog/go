package lib

import (
	"container/list"
	"fmt"
)

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimes(number int) *list.List {
	primes := list.New()
	for i := number; i >= 2; i-- {
		if isPrime(i) {
			fmt.Println(i)
			primes.PushFront(i)
		} else {
		}
	}
	return primes
}

func GetNthPrime(n int) int {
	count := 0
	number := 2
	for {
		//fmt.Println("number=", number)
		if isPrime(number) {
			count++
		}
		if count >= n {
			return number
		}
		number++
	}
	return 0
}
