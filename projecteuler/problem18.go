/*
https://projecteuler.net/problem=18
Maximum path sum I

By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3*
7* 4
2 4* 6
8 5 9* 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("problem 18")
	s := `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
	`
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	tree := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		tree[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			tmpInt, _ := strconv.Atoi(line[j])
			tree[i][j] = tmpInt
		}
	}

	fmt.Println("path: ", bruteForceMaxPathSum(tree))
}

func bruteForceMaxPathSum(tree [][]int) int {
	fmt.Println(tree)
	//largest := -1
	largest := recursivePaths(tree, 0, 0, 0, "")
	fmt.Println(largest)
	return largest
}

func recursivePaths(tree [][]int, x, y, pathSum int, pathStr string) int {
	//fmt.Printf("(%d, %d)\n", x, y)
	value := tree[x][y]
	//fmt.Println(value)
	pathStr += " + " + strconv.Itoa(value)
	pathSum += value
	max := 0
	if x < len(tree)-1 {
		left := recursivePaths(tree, x+1, y, pathSum, pathStr)
		right := recursivePaths(tree, x+1, y+1, pathSum, pathStr)
		if left > right {
			max = left
		} else {
			max = right
		}
	} else {
		//end of path
		fmt.Println("PATH STR=", pathStr)
		max = pathSum
	}
	return max
}
