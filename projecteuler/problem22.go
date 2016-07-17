/*
https://projecteuler.net/problem=22
Names scores

Using problem22.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("problem 22")
	names := getNames("./problem22.txt")
	sort.Strings(names)
	//fmt.Println(names)
	//fmt.Println(score("COLIN", 938))
	sum := 0
	for key, value := range names {
		sum += score(value, key+1)
	}
	fmt.Println("sum=", sum)
}

func score(s string, index int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		v, _ := utf8.DecodeRuneInString(c)
		tmpScore := int(v) - 64 //A==65
		//fmt.Println(c, tmpScore)
		sum += tmpScore
	}
	return sum * index
}

func getNames(path string) []string {
	names := make([]string, 0)
	lines, _ := getLines(path)
	for _, v := range lines {
		v = strings.Replace(v, "\"", "", -1)
		names = append(names, strings.Split(v, ",")...)
	}
	return names
}

func getLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
