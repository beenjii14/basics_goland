package main

func main() {
	data_types()
}

func data_types() {
	// Boolean
	var b bool = true
	println(b)

	// Integer
	var i int = 1
	println(i)

	// Floating Point
	var f float32 = 1.1
	println(f)

	// Complex
	var c complex64 = 1 + 2i
	println(c)

	// String
	var s string = "Hello World"
	println(s)

	// Array
	var a [3]int = [3]int{1, 2, 3}
	println(a)

	// Slice
	var sl []int = []int{1, 2, 3}
	println(sl)

	// Map
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	println(m)

	// Pointer
	var p *int = &i
	println(p)

	// Struct
	type Person struct {
		name string
		age  int
	}
	var person Person = Person{"John", 20}
	println(person)

	// Channel
	var ch chan int = make(chan int)
	println(ch)

	// Interface
	var inter interface{} = 1
	println(inter)
}
