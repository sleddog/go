/*
https://projecteuler.net/problem=26
Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("problem 26")
	//for i := 2; i < 100; i++ {
	//for i := 24; i < 25; i++ {
	//for i := 6; i < 7; i++ {
	for i := 95; i < 96; i++ {
		v := big.NewRat(int64(1), int64(i))
		c := recurringCycle(v)
		s := v.FloatString(100)
		fmt.Printf("1/%d: %d - %s\n", i, c, s)
	}

}

func recurringCycle(r *big.Rat) int {
	precision := 200
	s := strings.Split(r.FloatString(precision), ".")[1]
	//fmt.Println(s)
	lastChar := s[len(s)-1 : len(s)]
	if lastChar == "0" {
		return 0
	}
	//find the pattern
	a := strings.Split(s, s[0:1])
	if len(a) > precision {
		return 1
	}
	//try to grown the pattern
	//fmt.Println("len(s)=", len(s))
	for t := 0; t < len(s); t++ {
		for i := t; i < len(s)-t; i++ {
			//fmt.Println("t=", t)
			//fmt.Println("i=", i)
			subStr := s[t : t+i]
			//fmt.Println("subStr=", subStr)
			a = strings.Split(s, subStr)
			total := len(subStr) * len(a)
			if t > 0 {
				fmt.Println("t=", t)
				fmt.Println("len(subStr)=", len(subStr))
				fmt.Println("len(a)=", len(a))
				fmt.Println("total=", total)
				fmt.Println("precision=", precision)
				fmt.Println("subStr=", subStr)
				fmt.Println("len substring = ", len(subStr))
			}
			if len(a) > len(subStr) && total >= precision-t && (len(subStr) < precision/2) {
				fmt.Println(subStr)
				return i
			}
		}
	}
	return -1

}
