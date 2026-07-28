package experiments

import (
	"fmt"
	"time"
)

// gen creates a channel that sends numbers 1..n
func gen(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// square reads numbers from in channel, squares them, and sends to out
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

// double reads numbers from in channel, doubles them, and sends to out
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()
	return out
}

func Pipeline() {
	fmt.Println("Pipeline Example")

	// Stage 1: generate numbers
	nums := gen(5)

	// Stage 2: square the numbers
	squares := square(nums)

	// Stage 3: double the results
	doubles := double(squares)

	// Collect the results
	for v := range doubles {
		fmt.Println("Result:", v)
		time.Sleep(200 * time.Millisecond) // slow down for clarity
	}
}
