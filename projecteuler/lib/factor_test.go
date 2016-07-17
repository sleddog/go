package lib

import (
	"fmt"
	"testing"
)

func TestFactorFuncs(t *testing.T) {
	for i := 1; i < 100; i++ {
		factors := Factors(i)
		fmt.Println(factors)
		factors2 := Factors2(i)
		fmt.Println(factors2)

		if len(factors) != len(factors2) {
			t.Errorf("factor length must be the same")
		}
	}
}

func BenchmarkFactors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Factors(10000)
	}
}

func BenchmarkFactors2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Factors2(10000)
	}
}

func BenchmarkProperDivisors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProperDivisors(10000)
	}
}

func BenchmarkProperDivisors2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProperDivisors2(10000)
	}
}
