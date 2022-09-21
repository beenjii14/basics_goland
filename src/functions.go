package main

import "fmt"

// Function with a return value
func getGreeting(name string) string {
	return "Hello " + name
}

// Function with multiple return values
func getYears(age, age2 int, message string) (int, string) {
	return age + age2, message
}

func main() {
	fmt.Println(getGreeting("Benji"))

	// Multiple return values
	years, message := getYears(10, 20, "Hello")
	fmt.Println(years, message)

	// _ is a blank identifier
	_, message2 := getYears(10, 20, "Hello world")
	fmt.Println(message2)
}
