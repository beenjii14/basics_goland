package main

import "fmt"

func main() {
	message1 := "Hello, World!"
	message2 := "Hello, World! 2"

	// Println
	fmt.Println(message1, message2)

	// Printf
	name := "Benji"
	age := 28
	fmt.Printf("%s is %d years old\n", name, age)
	// %v is the default format
	fmt.Printf("%v is %v years old\n", name, age)

	// Sprintf
	message3 := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(message3)

	// %T prints the type of the variable
	fmt.Printf("age is of type %T\n", age)
}
