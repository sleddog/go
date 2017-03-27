/*
https://projecteuler.net/problem=58
Spiral primes

Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.

37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting is that 8 out of the 13 numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.

If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed. If this process is continued, what is the side length of the square spiral for which the ratio of primes along both diagonals first falls below 10%?
*/

package main

import (
	"./lib"
	"fmt"
)

func main() {
	fmt.Println("problem 58")
	targetRatio := 0.10
	n := 26239
	for {
		ratio := spiralPrimes(n)
		fmt.Println("length=", n, "ratio=", ratio)
		if ratio < targetRatio {
			fmt.Println("FOUND IT!")
			break
		}
		step := 2
		n = n + step
	}
}

func spiralPrimes(n int) float64 {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	y := n / 2
	x := n / 2
	grid[y][x] = 1

	//collect the diagonals (the corners of the latest grid)
	diagonals := []int{1} //start with center

	for size := 1; size < n; size = size + 2 {
		grid, y, x = doSpiral(y, x, size, grid)
		//printGrid(grid)
		topLeft := grid[y-size-1][x-size-1]
		topRight := grid[y-size-1][x]
		bottomLeft := grid[y][x-size-1]
		bottomRight := grid[y][x]
		diagonals = append(diagonals, topRight, topLeft, bottomLeft, bottomRight)
	}

	//fmt.Println(diagonals)
	return primeRatio(diagonals)
}

func primeRatio(numbers []int) float64 {
	primeCount := 0
	for _, n := range numbers {
		if lib.IsPrime(n) {
			primeCount++
		}
	}
	r := float64(primeCount) / float64(len(numbers))
	return r

}

//anticlockwise pattern: right 1, up (size), left (size+1), down (size+1), right (size+1)
func doSpiral(y, x, size int, grid [][]int) ([][]int, int, int) {
	v := grid[y][x]

	//right 1
	for i := 0; i < 1; i++ {
		x++
		v++
		grid[y][x] = v
	}

	//up (size)
	for i := 0; i < size; i++ {
		y--
		v++
		grid[y][x] = v
	}

	//left (size+1)
	for i := 0; i < size+1; i++ {
		x--
		v++
		grid[y][x] = v
	}

	//down (size+1)
	for i := 0; i < size+1; i++ {
		y++
		v++
		grid[y][x] = v
	}

	//right (size+1)
	for i := 0; i < size+1; i++ {
		x++
		v++
		grid[y][x] = v
	}

	return grid, y, x
}

func printGrid(grid [][]int) {
	fmt.Println("")
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%4d", grid[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}
