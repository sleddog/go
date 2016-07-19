/*
https://projecteuler.net/problem=36
Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 36")
	sum := 0
	for i := 0; i < 1000000; i++ {
		n := int64(i)
		s := strconv.FormatInt(n, 10)
		if isPalindrome(s) {
			b := strconv.FormatInt(n, 2)
			if isPalindrome(b) {
				fmt.Printf("%v and %v\n", s, b)
				sum += i
			}
		}
	}
	fmt.Println("sum=", sum)
}

func isPalindrome(s string) bool {
	r := reverseString(s)
	return s == r
}

func reverseString(val string) string {
	runes := []rune(val)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
