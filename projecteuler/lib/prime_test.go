package lib

import (
	"testing"
)

func TestGetNthPrime(t *testing.T) {
	//https://projecteuler.net/problem=7
	nthPrime := GetNthPrime(10001)
	if nthPrime != 104743 {
		t.Errorf("based on euler project 7, the 10001st prime should be 104743")
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrime(10000)
	}
}

func BenchmarkIsPrime2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrimeSlow(10000)
	}
}
