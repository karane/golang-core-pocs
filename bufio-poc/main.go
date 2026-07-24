package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "buffered.txt"

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Hello from bufio!\n")
	writer.WriteString("This is written using buffered I/O.\n")

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
	fmt.Println("Data written to file successfully with buffering!")

	file2, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	reader := bufio.NewReader(file2)

	fmt.Println("\nReading file line by line:")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// io.EOF means end of file
			if err.Error() == "EOF" {
				if len(line) > 0 {
					fmt.Print(line) // print last line if missing newline
				}
				break
			}
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Print(line)
	}

	fmt.Println("\nNow, type something (press Enter to stop):")
	stdinReader := bufio.NewReader(os.Stdin)
	input, _ := stdinReader.ReadString('\n')
	fmt.Println("You typed:", input)
}
