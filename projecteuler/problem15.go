/*
https://projecteuler.net/problem=15
Lattice paths

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 15")
	fmt.Println(numPaths(20))
}

func numPaths(gridSize int) int {
	return 2 * (1 + paths(gridSize, 1, 0))
}

func paths(gridSize, x, y int) int {
	//fmt.Printf("(%d, %d)\n", x, y)
	if x == gridSize {
		return 0
	}
	if y == gridSize {
		return 0
	}
	return 1 + paths(gridSize, x+1, y) + paths(gridSize, x, y+1)
}
