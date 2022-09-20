package main

import "fmt"

func main() {
	// Const variables are immutable
	const pi float64 = 3.14

	fmt.Println("The value of pi is:", pi)

	// types of variable declarations
	base := 12
	var height int = 14
	var area int

	fmt.Println("base:", base, "height:", height, "area:", area)

	// zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("int:", a, "float64:", b, "string:", c, "bool:", d)

}
