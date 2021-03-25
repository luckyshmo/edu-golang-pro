package main

import (
	"fmt"
	"io"
	"math"
)

type column struct {
	p []byte
}

func (c column) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	return 0, nil
}

func testInterface() {
	col := column{p: make([]byte, 5, 5)}
	col.Write(col.p)

	io.WriteString(col, "")

	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type square struct { //bcs of methods below square match Shape interface
	X float64
}
type circle struct { //bcs of methods below circle match Shape interface
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}
func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (s circle) Area() float64 {
	return s.R * s.R * math.Pi
}
func (s circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

func Calculate(x Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Is a circle!")
	}
	v, ok := x.(square) // ? intreface knows about realization
	if ok {
		fmt.Println("Is a square:", v)
	}
	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}
