package lib

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	strArr := strings.Split(s, "")
	sort.Strings(strArr)
	s = strings.Join(strArr, "")
	return s
}
