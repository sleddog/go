/*
https://projecteuler.net/problem=42
Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var nums map[int]bool

func main() {
	fmt.Println("problem 42")
	nums = make(map[int]bool)
	for i := 1; i <= 100000; i++ {
		n := triangleNum(i)
		//fmt.Printf("%d:%v\n", i, n)
		nums[n] = true

	}
	//fmt.Println(nums)
	//s := "SKY"
	//fmt.Println(s)
	//fmt.Println(isTriangleWord(s))

	//read word file
	words := getWords("./problem42.txt")
	count := 0
	for _, word := range words {
		if isTriangleWord(word) {
			count++
		}
	}
	fmt.Println("count =", count)
}

func isTriangleWord(s string) bool {
	sum := 0
	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		v, _ := utf8.DecodeRuneInString(c)
		tmpScore := int(v) - 64 //A==65
		//fmt.Printf("%v:%v\n", c, tmpScore)
		sum += tmpScore
	}
	//fmt.Println("sum=", sum)
	_, isInMap := nums[sum]
	return isInMap
}

func triangleNum(n int) int {
	return int(0.5 * float64(n) * (float64(n) + 1))
}

func getWords(path string) []string {
	words := make([]string, 0)
	lines, _ := getLines(path)
	for _, v := range lines {
		v = strings.Replace(v, "\"", "", -1)
		words = append(words, strings.Split(v, ",")...)
	}
	return words
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
