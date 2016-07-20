package lib

import (
//	"fmt"
)

func Permutations(input string) []string {
	r := len(input)
	entries := []string{}
	n := len(input)
	if r > n {
		return entries
	}

	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	//create the first entry
	entries = append(entries, input)

	entry := string(make([]byte, n))

	for n > 0 {
		i := r - 1
		for ; i >= 0; i -= 1 {
			cycles[i] -= 1
			if cycles[i] == 0 {
				index := indexes[i]
				for j := i; j < n-1; j += 1 {
					indexes[j] = indexes[j+1]
				}
				indexes[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indexes[i], indexes[n-j] = indexes[n-j], indexes[i]
				for k := i; k < r; k += 1 {
					//entry[k] = input[indexes[k]]
					entry = swapChar(entry, input[indexes[k]], k)
				}

				//we have an entry!
				//fmt.Println(entry)
				entries = append(entries, entry)
				break
			}
		}
		if i < 0 {
			return entries
		}
	}
	return entries
}

func swapChar(s string, swap byte, i int) string {
	return s[:i] + string(swap) + s[i+1:]
}
