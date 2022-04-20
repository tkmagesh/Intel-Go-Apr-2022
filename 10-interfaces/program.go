package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(x ShapeWithArea) {
	fmt.Printf("Area = %v\n", x.Area())
}

func PrintPerimeter( /*  */ ) {
	fmt.Printf("Perimeter = %v\n", x.Perimeter())
}

/*
Perimeter
	Circle = 2 * pi * r
	Rectangle = 2 * (height + width)
*/
func main() {
	c := Circle{Radius: 12}
	//fmt.Printf("Area = %v\n", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 10, Width: 12}
	//fmt.Printf("Area = %v\n", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
}
