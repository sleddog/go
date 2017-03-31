/*
https://projecteuler.net/problem=60
Prime pair sets

The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them in any order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
*/

package main

import (
	"./lib"
	"fmt"
	"sort"
	"strconv"
)

var primes []int

func main() {
	fmt.Println("problem 60")
	/*
		x := []int{3, 7, 109, 673}
		fmt.Println("allPrime=", testPrimes(x))
		fmt.Println("sum=", sum(x))
	*/

	max := 10000
	primes = lib.GetPrimesFromCache(max)
	//fmt.Println(primes)
	n := len(primes)
	fmt.Println(n)

	//two
	candidates := [][]int{}
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			p1 := primes[i]
			p2 := primes[j]
			x := []int{p1, p2}
			if testPrimes(x) {
				candidates = append(candidates, x)
			}
		}

	}

	//three
	candidates3 := [][]int{}
	for _, v := range candidates {
		for _, p := range primes {
			x := append(v, p)
			if testPrimes(x) {
				sort.Ints(x)
				if !inArray(candidates3, x) {
					candidates3 = append(candidates3, x)
				}
			}

		}
	}

	//four
	best := []int{}
	min := 999999999999999
	c4 := []string{}
	for _, v3 := range candidates3 {
		for _, p3 := range primes {
			x3 := append(v3, p3)
			if testPrimes(x3) {
				sort.Ints(x3)
				s := intArrToString(x3)
				if !stringInSlice(s, c4) {
					c4 = append(c4, s)
					sum, bestY := test5(x3)
					if sum <= min {
						min = sum
						best = bestY
					}
				}
			}

		}
	}
	fmt.Println("best=", best)
	fmt.Println("min=", min)

}

func sum(n []int) int {
	sum := 0
	for _, i := range n {
		sum += i
	}
	return sum

}

func testPrimes(n []int) bool {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			//A+B
			s1 := strconv.Itoa(n[i]) + strconv.Itoa(n[j])
			i1, _ := strconv.Atoi(s1)
			if !lib.IsPrime(i1) {
				return false
			}
			//B+A
			s2 := strconv.Itoa(n[j]) + strconv.Itoa(n[i])
			i2, _ := strconv.Atoi(s2)
			if !lib.IsPrime(i2) {
				return false
			}
		}
	}
	return true
}

func inArray(a [][]int, v []int) bool {
	for _, x := range a {
		if sliceEqual(v, x) {
			return true
		}
	}
	return false
}

func sliceEqual(a, b []int) bool {
	//fmt.Println("sliceEqual, a=", a, "b=", b)
	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func test5(x []int) (int, []int) {
	min := 999999999999999
	s := 0
	bestY := []int{}
	for _, p := range primes {
		y := append(x, p)
		if testPrimes(y) {
			fmt.Println("FOUND ONE")
			fmt.Println(y)
			s = sum(y)
			fmt.Println("sum=", s)
			if s <= min {
				min = s
				bestY = y
			}
		}

	}
	return min, bestY
}

func intArrToString(x []int) string {
	s := ""
	for _, v := range x {
		s += strconv.Itoa(v) + " "
	}
	return s
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
