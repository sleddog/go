/*
https://projecteuler.net/problem=19
Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import (
	"fmt"
)

type Day int

const (
	MONDAY Day = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

type Date struct {
	day       int
	month     int
	year      int
	dayOfWeek Day
}

func (d *Date) nextDate() Date {
	thirtyMonths := []int{4, 6, 9, 11}
	//thirtyOneMonths := []int{1, 3, 5, 7, 8, 12}
	dayCap := 0
	if d.month == 2 { //February
		//check if leap year
		if leapYear(d.year) {
			dayCap = 29
		} else {
			dayCap = 29
		}
	} else if intInSlice(d.month, thirtyMonths) {
		dayCap = 30
	} else {
		dayCap = 31
	}

	//increment
	var day, month, year int
	var dayOfWeek Day
	month = d.month
	year = d.year

	//change day
	if d.dayOfWeek == SUNDAY {
		dayOfWeek = MONDAY
	} else {
		dayOfWeek = d.dayOfWeek + 1
	}
	if d.day == dayCap {
		day = 1
		//change month
		if d.month == 12 {
			month = 1
			//change year
			year = d.year + 1
		} else {
			month = d.month + 1
		}
	} else {
		day = d.day + 1
	}
	fmt.Printf("day=%d, month=%d, year=%d, dayOfWeek=%d\n", day, month, year, dayOfWeek)
	nextDate := Date{day, month, year, dayOfWeek}
	return nextDate
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func leapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("problem 19")
	year := 1900
	month := 1
	day := 1
	dayOfWeek := MONDAY
	fmt.Println(year, month, day, dayOfWeek)
	date := Date{day, month, year, dayOfWeek}
	fmt.Println(date)
	//leap year test
	/*
		fmt.Println(2000, leapYear(2000))
		fmt.Println(1996, leapYear(1996))
		fmt.Println(1999, leapYear(1999))
	*/
	numSundays := 0
	for {
		date = date.nextDate()
		fmt.Println(date)
		if date.year >= 2001 {
			break
		}
		//now that we are in the range, we can count
		//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
		if date.year >= 1901 {
			if date.isSpecial() {
				numSundays++
			}
		}
	}
	fmt.Println("number of Sundays on first of the month: ", numSundays)
}

func (d *Date) isSpecial() bool {
	if d.day == 1 && d.dayOfWeek == SUNDAY {
		return true
	}
	return false
}
