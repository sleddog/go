/*
https://projecteuler.net/problem=28
Number spiral diagonals

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 28")
	fmt.Println(spiralDiagonalSum(1001))
}

func spiralDiagonalSum(n int) int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	//fmt.Println(grid)

	y := n / 2
	x := n / 2
	grid[y][x] = 1
	//printGrid(grid)

	for size := 1; size < n; size = size + 2 {
		grid, y, x = doSpiral(y, x, size, grid)
		//printGrid(grid)
		//fmt.Printf("%v, %v\n", y, x)
	}

	return sumDiagonals(grid)
}

func sumDiagonals(grid [][]int) int {
	sum := 0
	n := len(grid)
	for i := 0; i < n; i++ {
		sum += grid[i][i]
		sum += grid[n-i-1][i]
	}
	//subtract center since it was counted twice
	sum -= grid[n/2][n/2]
	return sum
}

/*
0: {0,1}     //right 1
1: {-size,0} //down (size)
2: {0,-size+1} //left (size+1)
3: {0,size+1}  //up (size+1)
4: {0,size+1}  //right (size+1)
*/

//pattern: right 1, down (size), left (size+1), up (size+1), right (size+1)
func doSpiral(y, x, size int, grid [][]int) ([][]int, int, int) {
	v := grid[y][x]

	//right 1
	for i := 0; i < 1; i++ {
		x++
		v++
		grid[y][x] = v
	}

	//down (size)
	for i := 0; i < size; i++ {
		y++
		v++
		grid[y][x] = v
	}

	//left (size+1)
	for i := 0; i < size+1; i++ {
		x--
		v++
		grid[y][x] = v
	}

	//up (size+1)
	for i := 0; i < size+1; i++ {
		y--
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
