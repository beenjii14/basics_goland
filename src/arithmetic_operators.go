package main

import "fmt"

func main() {
	arithmetic_operators()
}

func arithmetic_operators() {
	// Arithmetic operators
	// + - * / %
	x := 10
	y := 5

	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	fmt.Println("x % y =", x%y)

	// Increment and Decrement
	x++
	y--

	fmt.Println("x++ =", x)
	fmt.Println("y-- =", y)
}

func areaOfSquare() {
	// Area of Square
	// a^2
	a := 10
	fmt.Println("a^2 =", a*a)
}

func areaOfRectangle() {
	// Area of Rectangle
	// l * w
	l := 10
	w := 5
	fmt.Println("l * w =", l*w)
}

func areaOfTriangle() {
	// Area of Triangle
	// (b * h) / 2
	b := 10
	h := 5
	fmt.Println("(b * h) / 2 =", (b*h)/2)
}

func areaOfCircle() {
	// Area of Circle
	// πr^2
	pi := 3.14
	r := 10
	fmt.Println("πr^2 =", pi*float64(r*r))
}
