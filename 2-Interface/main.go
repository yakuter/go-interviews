package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Example of type assertion
	var shape Shape
	shape = Circle{Radius: 5}
	circle, isCircle := shape.(Circle)
	if isCircle {
		fmt.Println("Circle Area:", circle.Area())
	} else {
		fmt.Println("Not a Circle")
	}

	// Example of type switch
	shape = Rectangle{Width: 3, Height: 4}
	switch v := shape.(type) {
	case Circle:
		fmt.Println("Circle Area:", v.Area())
	case Rectangle:
		fmt.Println("Rectangle Area:", v.Area())
	default:
		fmt.Println("Unknown shape")
	}

	// Nil interface example
	var nilShape Shape
	if nilShape == nil {
		fmt.Println("Nil Shape")
	} else {
		fmt.Println("NotNil Shape")
	}
}
