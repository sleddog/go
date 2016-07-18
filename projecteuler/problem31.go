/*
https://projecteuler.net/problem=31
Coin Sums

In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 31")
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	numCoins := len(coins)
	fmt.Println("count = ", coinCount(coins, numCoins, 200))
}

func coinCount(coins []int, num, sumTarget int) int {
	if sumTarget == 0 {
		return 1 //no coins
	}

	if sumTarget < 0 {
		return 0 //can never create a negative coin
	}

	if num <= 0 && sumTarget >= 1 {
		return 0 //with no coins we cannot make a solution
	}

	//recursive calls
	return coinCount(coins, num-1, sumTarget) + coinCount(coins, num, sumTarget-coins[num-1])
}
