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

func TestIsNPandigitalInt(t *testing.T) {
	//problem 41 - 2143 is a 4-digit pandigital
	i := 2143
	if !IsNPandigitalInt(4, i) {
		t.Errorf("this number is 4-pandigital...")
	}
}
