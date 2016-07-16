package lib

import (
	"fmt"
	"testing"
)

func TestGetNthPrime(t *testing.T) {
	//https://projecteuler.net/problem=7
	nthPrime := GetNthPrime(10001)
	fmt.Println("10001st prime is", nthPrime)
	if nthPrime != 104743 {
		t.Errorf("based on euler project 7, the 10001st prime should be 104743")
	}
}
