package main

import (
	"fmt"

	"go-channels-poc/experiments"
)

func main() {
	fmt.Println("Go Channels Demo")
	fmt.Println("================")

	// Pick what you want to run:
	//experiments.BufferedChannel()
	//experiments.CloseChannel()
	//experiments.SelectStatement()
	//experiments.Pipeline()
	//experiments.FanInFanOut()
	//experiments.WorkerPool()
	experiments.Deadlock()
}
