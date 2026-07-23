package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, io.Reader and io.Writer!\n")

	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Read from strings.Reader:")
	fmt.Println(string(data))

	fmt.Println("\nWriting directly to os.Stdout using io.Writer:")
	writer := os.Stdout
	writer.Write([]byte("This is written using io.Writer!\n"))

	fmt.Println("\nCopying data from Reader to Writer (using io.Copy):")
	src := strings.NewReader("Copied from src to dst!\n")
	dst := os.Stdout
	n, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying:", err)
		return
	}
	fmt.Printf("Copied %d bytes.\n", n)

	fmt.Println("\nUsing bytes.Buffer as Reader and Writer:")
	var buf bytes.Buffer

	// Write to buffer
	buf.WriteString("First line in buffer.\n")
	buf.WriteString("Second line in buffer.\n")

	// Read back from buffer
	readData, err := io.ReadAll(&buf)
	if err != nil {
		fmt.Println("Error reading buffer:", err)
		return
	}
	fmt.Println(string(readData))

	fmt.Println("\nWriting and reading file with io.Reader and io.ReadAll:")
	os.WriteFile("example.txt", []byte("File I/O using io.Reader!\n"), 0644)
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(fileContent))
}
