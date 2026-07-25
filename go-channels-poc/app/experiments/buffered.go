package experiments

import "fmt"

func BufferedChannel() {
	fmt.Println("Buffered Channel Example")

	ch := make(chan int, 2) // buffer size 2

	ch <- 10
	ch <- 20

	fmt.Println("First:", <-ch)
	fmt.Println("Second:", <-ch)
}
