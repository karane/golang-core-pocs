package main

import (
	"fmt"
)

func PrintAll[T any](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

func Contains[T comparable](items []T, target T) bool {
	for _, v := range items {
		if v == target {
			return true
		}
	}
	return false
}

// Generic Set using comparable keys
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func main() {
	fmt.Println("PrintAll with any:")
	PrintAll([]int{1, 2, 3})
	PrintAll([]string{"go", "python", "rust"})
	PrintAll([]float64{1.1, 2.2})

	fmt.Println("\nContains:")
	fmt.Println(Contains([]int{1, 2, 3}, 2))       // true
	fmt.Println(Contains([]string{"a", "b"}, "c")) // false

	s := Set[string]{}
	s.Add("apple")
	s.Add("banana")

	fmt.Println("\nSet:")
	fmt.Println("Has apple?", s.Has("apple"))
	fmt.Println("Has orange?", s.Has("orange"))
}
