package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func main() {
	//https://gobyexample.com/hello-world
	fmt.Println("Hello Woodworker 2.0!")

	//https://gobyexample.com/values
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//https://gobyexample.com/variables
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	//https://gobyexample.com/constants
	fmt.Println(s)

	const n = 500000000

	const d2 = 3e20 / n
	fmt.Println(d2)

	fmt.Println(int64(d2))

	fmt.Println(math.Sin(n))

	//https://gobyexample.com/for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	//https://gobyexample.com/if-else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//https://gobyexample.com/switch
	i2 := 2
	fmt.Print("write ", i2, " as ")
	switch i2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	//https://gobyexample.com/arrays
	var a2 [5]int
	fmt.Println("emp:", a2)

	a2[4] = 100
	fmt.Println("set:", a2)
	fmt.Println("get:", a2[4])

	fmt.Println("len:", len(a2))

	b2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//https://gobyexample.com/slices
	s2 := make([]string, 3)
	fmt.Println("emp:", s2)

	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"
	fmt.Println("set:", s2)
	fmt.Println("get:", s2[2])

	fmt.Println("len:", len(s2))

	s2 = append(s2, "d")
	s2 = append(s2, "e", "f")
	fmt.Println("apd:", s2)

	c2 := make([]string, len(s2))
	copy(c2, s2)
	fmt.Println("cpy:", c2)

	l := s2[2:5]
	fmt.Println("sl1:", l)

	l = s2[:5]
	fmt.Println("sl2:", l)

	l = s2[2:]
	fmt.Println("sl3:", l)

	t2 := []string{"g", "h", "i"}
	fmt.Println("dcl:", t2)

	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD2)

	//https://gobyexample.com/maps
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n2)

	//https://gobyexample.com/range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
