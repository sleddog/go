package lib

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {
	s := "012"
	entries := Permutations(s)
	//012   021   102   120   201   210
	fmt.Println(entries)
	if len(entries) != 6 {
		t.Errorf("wrong number of permutations")
	}
}
