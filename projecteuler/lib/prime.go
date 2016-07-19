package lib

import (
	"container/list"
	"fmt"
)

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimes2(number int) []int {
	primes := []int{}
	for i := number; i >= 2; i-- {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
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

// A concurrent prime sieve

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func GetPrimesSieve(max int) []int {
	primes := []int{}
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.
	for {
		prime := <-ch
		//print(prime, "\n")
		if prime > max {
			break
		}
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}
