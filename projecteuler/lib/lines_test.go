//THIS LINE TESTS IS PART OF THE TEST!
package lib

import (
	"testing"
)

func TestGetLines(t *testing.T) {

	lines, _ := GetLines("./lines_test.go")
	expected := "//THIS LINE TESTS IS PART OF THE TEST!"
	if lines[0] != expected {
		t.Errorf("read lines failed")
	}
}
