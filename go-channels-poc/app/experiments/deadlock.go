package experiments

import "fmt"

func Deadlock() {
	fmt.Println("Deadlock Example")

	ch := make(chan int)

	// Uncomment to see the program hang
	//ch <- 1  // fatal error: all goroutines are asleep - deadlock!

	// Fix with goroutine:
	go func() { ch <- 1 }()
	val := <-ch
	fmt.Println("Received:", val)
}
