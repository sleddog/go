package lib

import (
	"container/list"
	"fmt"
)

func IsPrime(num int) bool {
	if num < 0 {
		return false //for problem27 (negative numbers are not prime)
	}
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
		if IsPrime(i) {
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
		if IsPrime(number) {
			count++
		}
		if count >= n {
			return number
		}
		number++
	}
	return 0
}
