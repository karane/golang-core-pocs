package main

import "fmt"

// Struct example
type Person struct {
	Name string
	Age  int
}

func main() {
	// Pointer basics
	fmt.Println("1. Pointer basics")
	x := 10
	fmt.Println("x before:", x)

	p := &x // & gives the address of x
	fmt.Println("Pointer p points to address:", p)
	fmt.Println("Value at pointer p:", *p) // * dereferences the pointer

	*p = 20 // change value via pointer
	fmt.Println("x after modifying via pointer:", x)

	fmt.Println("\n-------------------\n")

	// Passing by value vs reference
	fmt.Println("2. Passing by value vs reference")
	a := 5
	fmt.Println("Before doubling by value:", a)
	doubleValue(a) // passing by value
	fmt.Println("After doubleValue (by value):", a)

	fmt.Println("Before doubling by reference:", a)
	doubleReference(&a) // passing by reference
	fmt.Println("After doubleReference (by reference):", a)

	fmt.Println("\n-------------------\n")

	// Pointer with structs
	fmt.Println("3. Pointer with structs")
	person := Person{Name: "Alice", Age: 30}
	fmt.Println("Before birthday:", person)

	celebrateBirthday(&person) // passing struct pointer
	fmt.Println("After birthday:", person)
}

// Pass by value (copy of variable)
func doubleValue(n int) {
	n = n * 2
}

// Pass by reference (pointer)
func doubleReference(n *int) {
	*n = *n * 2
}

// Function that modifies struct via pointer
func celebrateBirthday(p *Person) {
	p.Age += 1
}
