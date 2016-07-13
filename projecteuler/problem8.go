/*
https://projecteuler.net/problem=8
Largest product in a series

The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450

Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
*/

package main

import (
	//	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("problem 8")
	adjacentLength := 13
	number := `
73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450
`
	fmt.Println(number)
	fmt.Println("len(number)=", len(number))
	//format the number a little bit
	number = strings.TrimSpace(number)
	number = strings.Replace(number, "\n", "", -1)
	fmt.Println(number)
	fmt.Println("len(number)=", len(number))

	//create a map of the position (int) and slice of int values (int, int, ... int)
	adjacentGroups := make(map[int][]int)
	fmt.Println(adjacentGroups)

	//create groups of adacent X digit numbers
	//currentLength := 0
	//currentGroupIndex := 0
	for i := 0; i < len(number); i++ {
		adjacentGroups[i] = make([]int, adjacentLength)
		//fmt.Println("i=", i)
		tmpInt, err := strconv.Atoi(number[i : i+1])
		if err != nil {
			fmt.Println(err)
			fmt.Println("SOMETHING IS BAD BRAH")
			break
		}
		//fmt.Println("tmpInt=", tmpInt)
		//add the number to all possible groups
		for x := 0; x < adjacentLength; x++ {
			if _, ok := adjacentGroups[i-x]; ok {
				adjacentGroups[i-x][x] = tmpInt
			}
		}
		//adjacentGroups[currentGroupIndex][currentLength] = tmpInt
		/*
			currentLength++
			if currentLength == adjacentLength {
				//create new group
				currentGroupIndex = i
				currentLength = 0
				adjacentGroups[currentGroupIndex] = make([]int, adjacentLength)
			}
		*/
	}
	//fmt.Println(adjacentGroups)
	largestProduct := 0
	largestProductIndex := -1
	for key, value := range adjacentGroups {
		//fmt.Println("Key:", key, "Value:", value)
		product := value[0]
		for i := 1; i < adjacentLength; i++ {
			product = product * value[i]
		}
		//fmt.Println("product=", product)
		if product >= largestProduct {
			largestProduct = product
			largestProductIndex = key
		}

	}
	fmt.Println(largestProduct)
	fmt.Println(largestProductIndex)
	/*
		fmt.Println(len(adjacentGroups))
		fmt.Println(adjacentGroups[0])
		fmt.Println(adjacentGroups[1])
		fmt.Println(adjacentGroups[2])
		fmt.Println(adjacentGroups[len(adjacentGroups)-4])
		fmt.Println(adjacentGroups[len(adjacentGroups)-3])
		fmt.Println(adjacentGroups[len(adjacentGroups)-2])
		fmt.Println(adjacentGroups[len(adjacentGroups)-1])
		fmt.Println(number)
	*/
	fmt.Println(adjacentGroups[largestProductIndex])
}
