/*
https://projecteuler.net/problem=55
Lychrel numbers

If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome. A number that never forms a palindrome through the reverse and add process is called a Lychrel number. Due to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a number is Lychrel until proven otherwise. In addition you are given that for every number below ten-thousand, it will either (i) become a palindrome in less than fifty iterations, or, (ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome. In fact, 10677 is the first number to be shown to require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations, 28-digits).

Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?

NOTE: Wording was modified slightly on 24 April 2007 to emphasise the theoretical nature of Lychrel numbers.
*/

package main

import (
	"fmt"
	"math/big"
	//	"strconv"
)

func main() {
	fmt.Println("problem 55")
	//n := 349
	//n := 196
	//n := 10677

	lychrelCount := 0
	for n := 1; n < 10000; n++ {
		bigN := big.NewInt(int64(n))
		iterations, palindrome := lychrel(bigN, 1)
		fmt.Printf("n: %d, iterations: %d, palindrome: %s\n", n, iterations, palindrome)
		if iterations >= 50 {
			//found a lychrel
			lychrelCount++
		}
	}
	fmt.Println("# Lychrel numbers: ", lychrelCount)
}

func lychrel(n *big.Int, depth int) (int, string) {
	//recursion base case
	if depth > 50 {
		return depth, "NOT FOUND"
	}

	//reverse
	nStr := n.String()
	reverseStr := reverseString(nStr)
	reverseBigInt := big.NewInt(0)
	_, ok := reverseBigInt.SetString(reverseStr, 10)
	if !ok {
		fmt.Println("Error...")
	}

	//add together
	sum := big.NewInt(0)
	sum.Add(n, reverseBigInt)
	sumStr := sum.String()

	if isPalindrome(sumStr) {
		return depth, sumStr
	} else {
		return lychrel(sum, depth+1)
	}
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
