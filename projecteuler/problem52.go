/*
https://projecteuler.net/problem=52
Permuted multiples

It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("problem 52")
	//b := sameDigits(12, 21)
	//fmt.Println(b)

	i := 125874
	i2 := i * 2
	fmt.Println(sameDigits(i, i2))

	for j := 1; j < 6; j++ {
		xStart := int(math.Pow10(j))
		xEnd := xStart*10 - 1
		fmt.Printf("xStart=%d, xEnd=%d \n", xStart, xEnd)
		for x := xStart; x < xEnd; x++ {
			mCount := 1
			for m := 2; m < 7; m++ {
				mx := m * x
				if sameDigits(x, mx) {
					mCount = m
				} else {
					break
				}
			}
			if mCount >= 3 {
				show(x)
				fmt.Println("mCount=", mCount)
				fmt.Println(x)
				if mCount == 6 {
					fmt.Println("finished")
					fmt.Println("mCount=", mCount)
					fmt.Println(x)
					return
				}
			}
		}
	}

}

func show(x int) {
	fmt.Println(x)
	for m := 2; m < 7; m++ {
		mx := m * x
		fmt.Println(mx)
	}
}

func sameDigits(i, j int) bool {
	s1 := strconv.Itoa(i)
	s1 = sortString(s1)

	s2 := strconv.Itoa(j)
	s2 = sortString(s2)

	return s1 == s2
}

func sortString(s string) string {
	strArr := strings.Split(s, "")
	sort.Strings(strArr)
	s = strings.Join(strArr, "")
	return s
}
