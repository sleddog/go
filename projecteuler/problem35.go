/*
https://projecteuler.net/problem=35
Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 35")
	//fmt.Println(isCircular(197))
	number := 1000000
	//number := 100
	primes := lib.GetPrimesSieve(number)
	count := 0
	for _, v := range primes {
		if isCircular(v) {
			count++
		}
	}
	fmt.Println("count=", count)

}

func isCircular(v int) bool {
	c := createCircles(v)
	for _, v := range c {
		if !lib.IsPrime(v) {
			return false
		}
	}
	return true
}

func createCircles(v int) []int {
	c := []int{}
	c = append(c, v)
	s := strconv.Itoa(v)
	l := len(s)
	for i := 1; i < l; i++ {
		firstPart := s[0:i]
		secondPart := s[i:l]
		tmpInt, _ := strconv.Atoi(secondPart + firstPart)
		c = append(c, tmpInt)
	}
	return c
}
