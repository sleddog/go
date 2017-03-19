/*
https://projecteuler.net/problem=51
Prime digit replacements

By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example having seven primes among the ten generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the first member of this family, is the smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an eight prime value family.
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("problem 51")
	//s := "*3"
	//s := "56**3"
	s := "*2*3*3"
	count, smallest := checkString(s)

	fmt.Println("count=", count)
	fmt.Println("smallest = ", smallest)

	//make perms
	//fmt.Println("starItUp=", starItUp(8))
}

func checkString(s string) (int, int) {

	count := 0
	primeFamily := []int{}
	for i := 0; i < 10; i++ {
		r := strconv.Itoa(i)
		x := strings.Replace(s, "*", r, -1)
		tmpInt, _ := strconv.Atoi(x)
		//fmt.Println(tmpInt)
		tmpStr := strconv.Itoa(tmpInt)
		if len(tmpStr) != len(s) {
			continue // kicks out replacement of 1st digit with 0
		}
		if lib.IsPrime(tmpInt) {
			count++
			primeFamily = append(primeFamily, tmpInt)
		}
	}
	fmt.Println("primeFamily=", primeFamily)
	if count > 0 {
		return count, primeFamily[0]
	}
	return count, -1
}

func starItUp(target int) []string {
	v := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "*"}

	s := ""
	count := 0
	smallest := 0

	strs := []string{}
	for _, v1 := range v {
		for _, v2 := range v {
			for _, v3 := range v {
				for _, v4 := range v {
					for _, v5 := range v {
						for _, v6 := range v {
							s = v1 + v2 + v3 + v4 + v5 + v6
							if strings.Contains(s, "*") {
								fmt.Println("checking s=", s)
								count, smallest = checkString(s)
								if count >= target {
									fmt.Println("count=", count)
									fmt.Println("smallest = ", smallest)
									strs = append(strs, s)
								}
							}
						}
					}
				}
			}
		}
	}
	return strs
}
