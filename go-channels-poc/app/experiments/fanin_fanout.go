package experiments

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer sends a sequence of numbers into a channel
func Producer(id, count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			val := rand.Intn(100)
			fmt.Printf("Producer %d -> %d\n", id, val)
			out <- val
			time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		}
		close(out)
	}()
	return out
}

// FanIn merges multiple input channels into one output channel
func FanIn(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			out <- v
		}
	}

	wg.Add(len(chs))
	for _, c := range chs {
		go output(c)
	}

	// Close out once all input channels are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// FanOut starts multiple workers that read from one channel
func FanOut(in <-chan int, numWorkers int) {
	var wg sync.WaitGroup

	worker := func(id int, c <-chan int) {
		defer wg.Done()
		for v := range c {
			fmt.Printf("Worker %d processing %d\n", id, v)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, in)
	}

	wg.Wait()
}

func FanInFanOut() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Fan-In / Fan-Out Example")

	// Fan-In: merge multiple producers
	p1 := Producer(1, 3)
	p2 := Producer(2, 3)
	p3 := Producer(3, 3)

	merged := FanIn(p1, p2, p3)

	// Fan-Out: distribute merged values to multiple workers
	FanOut(merged, 2)
}
