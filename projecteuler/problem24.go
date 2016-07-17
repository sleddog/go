/*
https://projecteuler.net/problem=24
Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

 012   021   102   120   201   210

 What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 24")
	//perm([]int{0, 1, 2}, 3, 4)
	perm([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 1000000)
}

//adapted from https://play.golang.org/p/JEgfXR2zSH
func perm(input []int, r int, nth int) {
	count := 0
	n := len(input)
	if r > n {
		return
	}

	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	//create the first entry
	entry := make([]int, r)
	for key, value := range indexes[:r] {
		entry[key] = input[value]
	}
	//fmt.Println(entry)
	count++

	for n > 0 {
		i := r - 1
		for ; i >= 0; i -= 1 {
			cycles[i] -= 1
			if cycles[i] == 0 {
				index := indexes[i]
				for j := i; j < n-1; j += 1 {
					indexes[j] = indexes[j+1]
				}
				indexes[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indexes[i], indexes[n-j] = indexes[n-j], indexes[i]
				for k := i; k < r; k += 1 {
					entry[k] = input[indexes[k]]
				}

				//we have an entry!
				//fmt.Println(entry)
				count++
				if count >= nth {
					fmt.Printf("%d-th permutation: %d\n", nth, entry)
					return
				}

				break
			}
		}
		if i < 0 {
			return
		}
	}
}
