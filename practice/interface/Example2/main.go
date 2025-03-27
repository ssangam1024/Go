package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func PrintArea(s Shape) {
	switch shape := s.(type) {
	case Circle:
		fmt.Printf("Shape is a Circle with radius %.2f and area %.2f\n", shape.Radius, shape.Area())
	case Rectangle:
		fmt.Printf("Shape is a Rectangle with width %.2f, height %.2f, and area %.2f\n", shape.Width, shape.Height, shape.Area())
	default:
		fmt.Println("Unknown shape")
	}
}
func main() {
	var x float64
	fmt.Print("Enter circle radius: ")
	fmt.Scan(&x)
	var y float64
	fmt.Print("Enter width of the rectangle: ")
	fmt.Scan(&y)
	var z float64
	fmt.Print("Enter height of the rectangle: ")
	fmt.Scan(&z)
	var s Shape = Circle{Radius: x}
	PrintArea(s)
	s = Rectangle{Width: y, Height: z}
	PrintArea(s)
}
