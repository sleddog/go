package lib

import (
	"strconv"
)

func IsPandigitalInt(n int) bool {
	s := strconv.Itoa(n)
	return IsPandigital(s)
}

func IsPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}

	m := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		tmpInt, _ := strconv.Atoi(s[i : i+1])
		if tmpInt == 0 {
			return false
		}
		if m[tmpInt] {
			return false
		}
		m[tmpInt] = true
	}
	if len(m) == 9 {
		return true
	} else {
		return false
	}
}
