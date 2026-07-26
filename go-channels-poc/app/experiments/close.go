package experiments

import "fmt"

func CloseChannel() {
	fmt.Println("Close Channel Example")

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println("Received:", v)
	}

	// Reading from a closed channel gives zero value
	val, ok := <-ch
	fmt.Println("After close:", val, ok) // 0 false
}
