package lib

import (
	"math/big"
)

func Factorial(n int) int {
	i := FactorialBigInt(n)
	return int(i.Int64())
}

func FactorialBigInt(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1.0)
	}
	product := big.NewInt(1)
	for i := 1; i <= n; i++ {
		product = product.Mul(product, big.NewInt(int64(i)))
	}
	return product
}
