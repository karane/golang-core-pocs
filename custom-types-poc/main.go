package main

import (
	"fmt"
	"strings"
)

// 1. Define a new type based on float64
type Celsius float64

// 2. Define a type alias for int
type MyIntAlias = int

// 3. Add a method to the new type
func (c Celsius) ToFahrenheit() float64 {
	return float64(c*9/5 + 32)
}

// 4. Define an interface
type Stringer interface {
	String() string
}

// 5. Make Celsius implement the Stringer interface
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

// 6. Another custom type (based on string)
type MyUpperString string

// 7. Attach a method so it implements Stringer
func (s MyUpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	// --- New type example ---
	var temp Celsius = 25
	fmt.Println("Temperature in Celsius:", temp)
	fmt.Println("Temperature in Fahrenheit:", temp.ToFahrenheit())

	// --- Type alias example ---
	var x int = 10
	var y MyIntAlias = 20
	sum := x + y // works without conversion
	fmt.Println("Sum using type alias:", sum)

	// --- Custom type vs alias difference ---
	var c1 Celsius = 30
	// var z int = c1   // compile error: different types
	var z int = int(c1) // explicit conversion required
	fmt.Println("Converted Celsius to int:", z)

	// For alias, no conversion is needed
	var a MyIntAlias = 50
	var b int = a // same type
	fmt.Println("Assigned alias to int directly:", b)

	// --- Custom type implementing interface ---
	var s Stringer

	s = temp // Celsius implements Stringer
	fmt.Println("Celsius via Stringer:", s.String())

	s = MyUpperString("golang interfaces")
	fmt.Println("MyUpperString via Stringer:", s.String())
}
