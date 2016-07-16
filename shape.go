package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

type point struct {
	x int
	y int
}

var img = image.NewRGBA(image.Rect(0, 0, 500, 500))
var col color.Color

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a vertical line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}

func Polygon(points ...point) {
	l := len(points)
	if l < 2 {
		fmt.Println("not enough points specified")
		return
	}
	for i := 0; i < l-1; i++ {
		Line(points[i], points[i+1])
	}
	if l > 2 {
		Line(points[l-1], points[0])
	}
	/*
		p1 := points[0]
		p2 := points[1]
		p3 := points[2]
		p4 := points[3]
		Line(p1, p2)
		Line(p2, p3)
		Line(p3, p1)
		Line(p4, p1)
	*/
	/*
		for _, p := range points {
			if(
			fmt.Println("x = ", p.x)
			fmt.Println("y = ", p.y)
			Line(p1, p2)
		}
	*/
}

func Line(p1, p2 point) {
	fmt.Println("p1=", p1.x, p1.y)
	fmt.Println("p2=", p2.x, p2.y)

	//draw a line from p1 to p2
	x := p1.x
	y := p1.y

	//need floats to fake a skewed line
	xf := float64(x)
	yf := float64(y)

	xStep := 0
	yStep := 0

	//determine what to increment x by
	if p1.x < p2.x {
		xStep = 1
	} else if p1.x == p2.x {
		xStep = 0
	} else {
		xStep = -1
	}
	//determine what to increment x by
	if p1.y < p2.y {
		yStep = 1
	} else if p1.y == p2.y {
		yStep = 0
	} else {
		yStep = -1
	}

	fmt.Println("xStep=", xStep)
	fmt.Println("yStep=", yStep)

	//based on the distance between the points, set the xf and yf steps
	xdiff := 0.0
	if xStep == 1 {
		xdiff = float64(p2.x) - float64(p1.x)
	} else if xStep == -1 {
		xdiff = float64(p1.x) - float64(p2.x)
	}
	fmt.Println("xdiff = ", xdiff)

	ydiff := 0.0
	if yStep == 1 {
		ydiff = float64(p2.y) - float64(p1.y)
	} else if yStep == -1 {
		ydiff = float64(p1.y) - float64(p2.y)
	}
	fmt.Println("ydiff = ", ydiff)

	//based on whatever is largest diff, use that one as the 1-stepper
	var xfStep, yfStep float64
	if xdiff > ydiff {
		xfStep = 1.0
		yfStep = ydiff / xdiff
	} else {
		xfStep = xdiff / ydiff
		yfStep = 1.0
	}
	//now set the sign correctly
	xfStep = xfStep * float64(xStep)
	yfStep = yfStep * float64(yStep)
	fmt.Println("xfStep = ", xfStep)
	fmt.Println("yfStep = ", yfStep)

	for {
		//fmt.Println("x=", x)
		//fmt.Println("y=", y)
		img.Set(x, y, col)

		//x = x + xStep
		xf = xf + xfStep
		fmt.Println("xf=", xf)
		//get the decimal part of xf and yf
		intx := int(xf)
		decx := xf - float64(intx)
		//fmt.Println("xf = ", xf)
		//fmt.Println("intx = ", intx)
		fmt.Println("decx = ", decx)
		if decx > 0.5 {
			x = intx + 1
		} else {
			x = intx
		}

		//y = y + yStep
		yf = yf + yfStep
		fmt.Println("yf=", yf)
		inty := int(yf)
		decy := yf - float64(inty)
		//fmt.Println("yf = ", yf)
		//fmt.Println("inty = ", inty)
		fmt.Println("decy = ", decy)
		if decy > 0.5 {
			y = inty + 1
		} else {
			y = inty
		}

		//if xStep == 1 && x > p2.x {
		if xStep == 1 && xf > float64(p2.x) {
			fmt.Println("BREAK 1")
			break
		}
		//if xStep == -1 && x < p2.x {
		if xStep == -1 && xf < float64(p2.x) {
			fmt.Println("BREAK 2")
			break
		}
		//if yStep == 1 && y > p2.y {
		if yStep == 1 && yf > float64(p2.y) {
			fmt.Println("BREAK 3")
			break
		}
		//if yStep == -1 && y < p2.y {
		if yStep == -1 && yf < float64(p2.y) {
			fmt.Println("BREAK 4")
			break
		}
	}
}

func randomPolygon(numSides, maxSize int) {
	fmt.Printf("generate a random polygon with %d sides\n", numSides)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	points := make([]point, numSides)
	for i := 0; i < numSides; i++ {
		randX := r1.Intn(maxSize)
		randY := r1.Intn(maxSize)
		points[i] = point{randX, randY}
	}
	//Polygon(point{5, 5}, point{15, 105}, point{105, 155}, point{55, 64}, point{45, 78})
	Polygon(points...)
}

func regularPolygon(numSides, radius int) {
	fmt.Printf("generate a regular polygon with %d sides, %d radius\n", numSides, radius)
	centerX := 100
	centerY := 100
	points := make([]point, numSides)
	for i := 0; i < numSides; i++ {
		x := float64(centerX) + float64(radius)*math.Cos(2.0*math.Pi*float64(i)/float64(numSides))
		y := float64(centerY) + float64(radius)*math.Sin(2.0*math.Pi*float64(i)/float64(numSides))
		fmt.Printf("(%f, %f)\n", x, y)
		intX := int(x)
		intY := int(y)
		points[i] = point{intX, intY}
	}
	//Polygon(point{5, 5}, point{15, 105}, point{105, 155}, point{55, 64}, point{45, 78})
	Polygon(points...)
}

func main() {

	//collect arguments
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	//fmt.Println(arg)

	//col = color.RGBA{255, 0, 0, 255} // Red
	//HLine(10, 20, 80)
	//col = color.RGBA{0, 255, 0, 255} // Green
	//Rect(10, 10, 80, 50)

	//play with structs
	col = color.RGBA{0, 0, 255, 255} // Blue

	//box
	//Polygon(point{1, 1}, point{1, 100}, point{100, 100}, point{100, 1})

	//triangle
	col = color.RGBA{0, 255, 255, 255} // cyan
	//Polygon(point{5, 5}, point{15, 105}, point{105, 155})

	//line?
	//col = color.RGBA{255, 255, 255, 255} // white
	//Polygon(point{1, 1}, point{2, 100})

	//col = color.RGBA{255, 255, 255, 255} // white
	//Polygon(point{1, 1}, point{100, 100})

	//col = color.RGBA{0, 255, 0, 255} // Green
	//col = color.RGBA{230, 255, 0, 255} // Green
	//randomPolygon(62, 300)
	//col = color.RGBA{255, 255, 255, 255} // white
	//randomPolygon(8, 300)

	col = color.RGBA{0, 0, 0, 255} // black
	//regularPolygon(6, 80)

	//	regularPolygon(5, 80)
	//	regularPolygon(5, 50)
	regularPolygon(8, 100)

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
