//https://projecteuler.net/problem=4
//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("problem 4")
	fmt.Println(largestPalindromeFrom3DigitNumbers())
}

func largestPalindromeFrom3DigitNumbers() int {
	largest := -1
	//first lets try to print out all of the 3 digit numbers
	l := all3DigitNumbers()
	l2 := all3DigitNumbers()
	for e := l.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
		intVal1, ok := e.Value.(int)
		if !ok {
			fmt.Println("SOMETHINGS NOT RIGHT")
		}
		for e2 := l2.Front(); e2 != nil; e2 = e2.Next() {
			//fmt.Print(e2.Value, " ")

			intVal2, ok2 := e2.Value.(int)
			if !ok2 {
				fmt.Println("SOMETHINGS NOT RIGHT")
			}

			//multiply together
			product := intVal1 * intVal2
			//fmt.Printf("%d * %d = %d\n", intVal1, intVal2, product)
			if isPalindrome(product) {
				if product > largest {
					largest = product
					fmt.Printf("%d * %d = %d\n", intVal1, intVal2, product)
				}
			}
		}
	}
	return largest
}

func all3DigitNumbers() *list.List {
	l := list.New()
	for i := 100; i < 1000; i++ {
		l.PushFront(i)
	}
	return l
}

func isPalindrome(val int) bool {
	s := strconv.Itoa(val)
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
