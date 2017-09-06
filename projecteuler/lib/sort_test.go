package lib

import (
	"testing"
)

func TestSortString(t *testing.T) {
	s := "bdac"
	sorted := SortString(s)
	if sorted != "abcd" {
		t.Errorf("SortString is wrong...")
	}
}
