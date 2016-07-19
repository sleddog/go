package lib

import (
	"testing"
)

func TestIsPandigital(t *testing.T) {
	i := "192384576"
	if !IsPandigital(i) {
		t.Errorf("this number is pandigital...")
	}
}

func TestIsPandigitalInt(t *testing.T) {
	i := 192384576
	if !IsPandigitalInt(i) {
		t.Errorf("this number is pandigital...")
	}
}
