package main

import (
	"fmt"
	"math"
	"./myInterface" // relative import path - bad practice
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

// Area() is a behavior defined in Shape intf
// Here, square type implements this Area behavior
// using the Type method syntax. A type method is a
// function with 'receiver' as an argument
func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (s circle) Area() float64 {
	return s.R * s.R * math.Pi
}

// Try to remove this. Compiler will complain that
// 'circle' doesn't is not a Shape, i.e., Perimeter()
// is undefined
func (s circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

// Here is the cool part. You don't need to have
// a Calculate function each for {square, circle, rect}
// We know that object that implements the Shape intf
// have Area() and Perimeter() defined.
func Calculate(x myInterface.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Is a circle!")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("Is a square:", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)
}
