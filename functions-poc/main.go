package main

import (
	"fmt"
)

func main() {
	// 1. Function declaration and calling
	sum := add(3, 5)
	fmt.Println("1. Function declaration - Sum:", sum)
	fmt.Println()

	// 2. Function with multiple return values
	a, b := swap("hello", "world")
	fmt.Println("2. Multiple return values - Swapped:", a, b)
	fmt.Println()

	// 3. Function with named return values
	result := divide(10, 3)
	fmt.Println("3. Named return values - Division result:", result)
	fmt.Println()

	// 4. Variadic function
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("4. Variadic function - Sum of all numbers:", total)
	fmt.Println()

	// 5a. Anonymous function immediately invoked
	func() {
		fmt.Println("5a. Anonymous function - Immediately invoked")
		fmt.Println()

	}()

	// 5b. Anonymous function assigned to a variable and called later
	greet := func(name string) {
		fmt.Println("5b. Anonymous function - Hello", name)
		fmt.Println()
	}
	greet("Go Developer")

	// 6a. Simple closure
	counter := createCounter()
	fmt.Println("6a. Closure - counter call 1:", counter())
	fmt.Println("6a. Closure - counter call 2:", counter())
	fmt.Println("6a. Closure - counter call 3:", counter())
	fmt.Println()

	// 6b. Idiomatic closure combined with anonymous function
	closure := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}() // immediately invoked to return the closure

	fmt.Println("6b. Closure via anonymous function - call 1:", closure())
	fmt.Println("6b. Closure via anonymous function - call 2:", closure())
	fmt.Println("6b. Closure via anonymous function - call 3:", closure())
	fmt.Println()

	// 6c. Closure capturing outer variable
	adder := func(x int) func(int) int {
		return func(y int) int {
			return x + y // x is captured
		}
	}(10) // x = 10

	fmt.Println("6c. Closure capturing outer value:", adder(5))  // 15
	fmt.Println("6c. Closure capturing outer value:", adder(20)) // 30
	fmt.Println()
}

// 1. Function declaration
func add(x int, y int) int {
	return x + y
}

// 2. Multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// 3. Named return values
func divide(x, y float64) (result float64) {
	result = x / y
	return // implicit return of 'result'
}

// 4. Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// 6a. Simple closure
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
