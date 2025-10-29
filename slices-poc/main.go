package main

import "fmt"

func main() {
	// 1. Arrays: declaration and initialization
	var arr1 [3]int              // declaration, zero values
	arr2 := [3]int{1, 2, 3}      // declaration + initialization
	arr3 := [...]int{4, 5, 6, 7} // length inferred

	fmt.Println("1. Arrays:")
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println()

	// 2. Slices: creation, append, copy, slicing
	slice1 := []int{10, 20, 30}     // slice literal
	slice2 := make([]int, 3, 5)     // length 3, capacity 5
	slice3 := slice1[1:3]           // slicing (from index 1 to 2)
	slice2 = append(slice2, 40, 50) // append elements

	copySlice := make([]int, len(slice1))
	copy(copySlice, slice1) // copy elements

	fmt.Println("2. Slices:")
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2 (after append):", slice2)
	fmt.Println("slice3 (sliced from slice1):", slice3)
	fmt.Println("copySlice (copied from slice1):", copySlice)
	fmt.Println()

	// 3. Slice vs Array differences
	fmt.Println("3. Slice vs Array differences:")
	fmt.Println("len(arr2):", len(arr2), "cap(arr2):", cap(arr2))
	fmt.Println("len(slice1):", len(slice1), "cap(slice1):", cap(slice1))
	fmt.Println("Slices are dynamic, arrays are fixed.")
	fmt.Println()

	// 4. Nil slices
	var nilSlice []int
	emptySlice := []int{}
	fmt.Println("4. Nil vs empty slice:")
	fmt.Println("nilSlice:", nilSlice, "len:", len(nilSlice), "nil?", nilSlice == nil)
	fmt.Println("emptySlice:", emptySlice, "len:", len(emptySlice), "nil?", emptySlice == nil)
	fmt.Println()

	// 5. Reslicing and underlying array sharing
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4] // slice from arr
	fmt.Println("5. Reslicing and array sharing:")
	fmt.Println("Original arr:", arr)
	fmt.Println("Slice s:", s)

	s[0] = 100 // modifies arr as well
	fmt.Println("Modified slice s:", s)
	fmt.Println("Array after slice modification:", arr)

	s = s[:cap(s)] // reslicing up to capacity
	fmt.Println("Resliced s up to capacity:", s)
	fmt.Println()

	// 6. Multidimensional slices (jagged)
	multiSlice := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}

	fmt.Println("6. Multidimensional slices:")
	for i, row := range multiSlice {
		fmt.Printf("Row %d: %v\n", i, row)
	}
	fmt.Println()

	// 7. Append and capacity growth demonstration
	growSlice := []int{1, 2}
	fmt.Println("7. Append and capacity growth:")
	fmt.Printf("Initial: len=%d, cap=%d, slice=%v\n", len(growSlice), cap(growSlice), growSlice)
	for i := 3; i <= 10; i++ {
		growSlice = append(growSlice, i)
		fmt.Printf("After append %d: len=%d, cap=%d, slice=%v\n", i, len(growSlice), cap(growSlice), growSlice)
	}
}
