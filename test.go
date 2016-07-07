package main

import (
	"fmt"
	//"math"
	"strconv"
)

func main() {
	/*
		x := 0.1
		fmt.Println("x=", x)

		for i := 0.0; i < 2.0; i += 0.1 {
			fmt.Println("i=", i)
			if i >= 0.5 {
				fmt.Println("go up", math.Ceil(i))
			} else {
				fmt.Println("go down", math.Floor(i))
			}
		}
		fmt.Println("floor=", math.Floor(x))
		fmt.Println("ceil=", math.Ceil(x))
	*/

	i := 10001
	fmt.Println(isPalindrome(i))
	i = 20001
	fmt.Println(isPalindrome(i))
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
