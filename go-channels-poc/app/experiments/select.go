package experiments

import (
	"fmt"
	"time"
)

func SelectStatement() {
	fmt.Println("Select Statement Example")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}
