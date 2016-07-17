/*
https://projecteuler.net/problem=17
Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("problem 17")
	totalCount := 0
	for i := 1; i <= 1000; i++ {
		word := numberToWord(i)
		length := measureString(word)
		totalCount += length
		fmt.Printf("%d=%s (%d)\n", i, word, length)
	}
	fmt.Println("total=", totalCount)
	/*
		fmt.Println(numberToWord(100))
		fmt.Println(numberToWord(101))
		fmt.Println(numberToWord(104))
		fmt.Println(numberToWord(110))
		fmt.Println(numberToWord(113))
		fmt.Println(numberToWord(119))
		fmt.Println(numberToWord(120))
	*/
}

func numberToWord(num int) string {
	if num >= 1000 {
		return "one thousand"
	}
	if num <= 20 {
		return simpleNum(num)
	}
	hundreds := num / 100
	str := ""
	if hundreds > 0 {
		str += simpleNum(hundreds) + " hundred"
	}
	num = num - hundreds*100
	tens := num / 10
	if tens >= 2 {
		if len(str) > 0 {
			str += " and "
		}
		str += simpleTens(tens * 10)
		remainder := num % 10
		if remainder > 0 {
			str += "-" + simpleNum(remainder)
		}
	} else {
		tmpStr := simpleNum(num)
		if len(tmpStr) > 0 {
			str += " and " + tmpStr
		}
	}
	return str
}

func simpleNum(num int) string {
	switch num {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	case 20:
		return "twenty"
	}
	return ""
}

func simpleTens(num int) string {
	switch num {
	case 20:
		return "twenty"
	case 30:
		return "thirty"
	case 40:
		return "forty"
	case 50:
		return "fifty"
	case 60:
		return "sixty"
	case 70:
		return "seventy"
	case 80:
		return "eighty"
	case 90:
		return "ninety"
	}
	return ""
}

func measureString(num string) int {
	num = strings.TrimSpace(num)
	num = strings.Replace(num, "-", "", -1)
	num = strings.Replace(num, " ", "", -1)
	return len(num)
}
