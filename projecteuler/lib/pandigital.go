package lib

import (
	"strconv"
)

func IsNPandigitalInt(n, v int) bool {
	s := strconv.Itoa(v)
	return IsNPandigital(n, s)
}

func IsNPandigital(n int, s string) bool {
	if len(s) != n {
		return false
	}
	m := make([]int, n)
	for i := 0; i < len(s); i++ {
		tmpInt, _ := strconv.Atoi(s[i : i+1])
		if tmpInt == 0 || tmpInt > n {
			return false
		}
		if m[tmpInt-1] != 0 {
			return false
		}
		m[tmpInt-1] = 1
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	if sum == n {
		return true
	} else {
		return false
	}
}

func IsPandigitalInt(n int) bool {
	return IsNPandigitalInt(9, n)
}

func IsPandigital(s string) bool {
	return IsNPandigital(9, s)
}

func IsPandigital10(s string) bool {
	n := 10
	if len(s) != n {
		return false
	}
	m := make([]int, 10)
	for i := 0; i < len(s); i++ {
		tmpInt, _ := strconv.Atoi(s[i : i+1])
		if m[tmpInt] != 0 {
			return false
		}
		m[tmpInt] = 1
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	if sum == n {
		return true
	} else {
		return false
	}
}
