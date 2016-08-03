/*
https://projecteuler.net/problem=49
Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?

*/

package main

import (
	"./lib"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("problem 49")
	p := lib.GetPrimesSieve(10000)
	m := make(map[string][]int)
	for _, v := range p {
		s := strconv.Itoa(v)
		if len(s) == 4 {
			m[s] = []int{v}
		}
	}
	for x, _ := range m {
		perms := lib.Permutations(x)
		for _, i := range perms {
			//convert to int then string again
			tmpInt, _ := strconv.Atoi(i)
			tmpStr := strconv.Itoa(tmpInt)
			if len(tmpStr) == 4 {
				if lib.IsPrime(tmpInt) {
					if !isInList(tmpInt, m[x]) {
						m[x] = append(m[x], tmpInt)
					}
				}
			}
		}
	}
	for q, w := range m {
		if len(w) == 3 {
			sort.Ints(w)
			if testList(w) {
				fmt.Println(q, w)
			}
		}
	}
}

func isInList(v int, l []int) bool {
	for _, n := range l {
		if n == v {
			return true
		}
	}
	return false
}

func testList(l []int) bool {
	diff1 := l[1] - l[0]
	diff2 := l[2] - l[1]
	return diff1 == diff2
}
