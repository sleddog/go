/*
https://projecteuler.net/problem=62
Cubic permutations

The cube, 41063625 (345^3), can be permuted to produce two other cubes: 56623104 (384^3) and 66430125 (405^3). In fact, 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are cube.
*/

package main

import (
	"./lib"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("problem 62")
	//fmt.Println(345, "^ 3 = ", cube(big.NewInt(int64(345))))
	//fmt.Println(384, "^ 3 = ", cube(big.NewInt(int64(384))))
	//fmt.Println(405, "^ 3 = ", cube(big.NewInt(int64(405))))

	// pre generate a range of cubes
	cubes := map[int]string{}
	for n := 100; n < 10000; n++ {
		c := cube(big.NewInt(int64(n)))
		cubes[n] = c.String()
	}

	// now lets do some intelligent pre-guessing by sorting the
	// cube string values and checking only large groups
	groups := findGroups(cubes)
	for sorted, cubes := range groups {
		if len(cubes) >= 5 {
			fmt.Println("sorted=", sorted, "length = ", len(cubes), "cubes = ", cubes)
			for _, cube := range cubes {
				fmt.Println("Testing ", cube)
			}
		}
	}
	// look at the output and find the smallest value from output above ^^
}

func cube(x *big.Int) *big.Int {
	value := new(big.Int).Exp(x, big.NewInt(int64(3)), nil)
	return value
}

func findGroups(cubes map[int]string) map[string][]string {
	groups := map[string][]string{}

	for _, cube := range cubes {
		//sort the cube
		sorted := lib.SortString(cube)
		//put the cube in the array IN in the map key' on this sorted value
		groups[sorted] = append(groups[sorted], cube)
	}
	return groups
}
