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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(x ShapeWithArea) {
	fmt.Printf("Area = %v\n", x.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %v\n", sp.Perimeter())
}

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
