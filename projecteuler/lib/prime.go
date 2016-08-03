package lib

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
)

func IsPrimeSlow(num int) bool {
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

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
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

func GetPrimesFromCache(max int) []int {
	tmpStr := strconv.Itoa(max)
	fileName := "./lib/primes" + tmpStr + ".gob"
	if _, err := os.Stat(fileName); err == nil {
		return readFile(fileName)
	} else {
		p := GetPrimesSieve(max)
		writeFile(fileName, p)
		return p
	}
}

func writeFile(fileName string, p []int) {
	// create a file
	dataFile, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(p)

	dataFile.Close()
}

func readFile(fileName string) []int {
	var data []int

	// open data file
	dataFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFile.Close()
	return data
}
