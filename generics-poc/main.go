package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// --- 3. Custom constraint: “Number” ---
// Number restricts T to int/int64/float32/float64 (or named types built on
// them, thanks to ~), so Sum can use += without runtime type checks.
type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

// Generic sum using custom constraint
func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

func main() {
	fmt.Println("Max int:", Max(10, 20))
	fmt.Println("Max float:", Max(3.14, 2.71))

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	v, _ := intStack.Pop()
	fmt.Println("Popped from int stack:", v)

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	s, _ := stringStack.Pop()
	fmt.Println("Popped from string stack:", s)

	fmt.Println("Sum ints:", Sum([]int{1, 2, 3, 4}))
	fmt.Println("Sum floats:", Sum([]float64{1.5, 2.5, 3.5}))
}
