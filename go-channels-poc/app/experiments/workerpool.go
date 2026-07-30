package experiments

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		results <- fmt.Sprintf("Worker %d processed task %d", id, task.ID)
	}
}

func WorkerPool() {
	rand.Seed(time.Now().UnixNano())

	numWorkers := 3
	numTasks := 5

	tasks := make(chan Task, numTasks)
	results := make(chan string, numTasks)

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
	fmt.Println("All tasks done ✅")
}
	