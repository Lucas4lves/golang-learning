package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Interfaces")

	n := NetworkingError{
		Networking: false,
		Hardware:   false,
	}

	fmt.Println(n)
}

type NetworkingError struct {
	Networking bool
	Hardware   bool
}

func (n NetworkingError) Error() string {
	if n.Networking {
		return "Networking error"
	} else if n.Hardware {
		return "Hardware error"
	}

	return "Another problem"

}

type GeometricForm interface {
	CalcArea() float64
}

func PrintArea(g GeometricForm) float64 {
	return g.CalcArea()
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) CalcArea() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (r Rectangle) CalcArea() float64 {
	return r.width * r.height
}
