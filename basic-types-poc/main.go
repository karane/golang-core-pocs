package main

import (
	"fmt"
)

func main() {
	// 1. Variables
	fmt.Println("1. Variables")
	// Single variable with explicit type
	var age int = 30
	fmt.Println("Age:", age)

	// Multiple variables with explicit types
	var x, y, z int = 1, 2, 3
	fmt.Println("x, y, z:", x, y, z)

	// Short declaration with type inference
	name := "Karane"
	country, city := "Brazil", "São Paulo"
	fmt.Println("Name:", name)
	fmt.Println("Country, City:", country, city)
	fmt.Println()

	// 2. Constants
	fmt.Println("2. Constants")
	// Untyped constant
	const pi = 3.14159
	// Typed constant
	const maxScore int = 100
	fmt.Println("Pi (untyped):", pi)
	fmt.Println("Max Score (typed int):", maxScore)
	fmt.Println()

	// 3. Basic types
	fmt.Println("3. Basic types")
	var height float64 = 1.75      // float64
	var smallFloat float32 = 3.14  // float32
	var isStudent bool = true      // boolean
	var greeting string = "Hello!" // string
	fmt.Println("Height:", height)
	fmt.Println("Small float:", smallFloat)
	fmt.Println("Is student?", isStudent)
	fmt.Println("Greeting:", greeting)
	fmt.Println()

	// 4. Type conversion
	fmt.Println("4. Type conversion")
	var score int = 95
	var scoreFloat float64 = float64(score) // int → float64
	fmt.Println("Score as float:", scoreFloat)

	// float32 → float64
	var bigFloat float64 = float64(smallFloat)
	fmt.Println("Converted float32 to float64:", bigFloat)

	// int → int64
	var largeInt int64 = int64(score)
	fmt.Println("Score as int64:", largeInt)
	fmt.Println()

	// 5. Zero values
	fmt.Println("5. Zero values")
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroPointer *int // nil pointer
	fmt.Println("Zero int:", zeroInt)
	fmt.Println("Zero float:", zeroFloat)
	fmt.Println("Zero bool:", zeroBool)
	fmt.Println("Zero string:", zeroString)
	fmt.Println("Zero pointer:", zeroPointer)
}
