package main

import "strings"

// Reverse returns the reversed version of a string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords returns the number of words in a string.
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	println(Reverse("golang"))
	println(CountWords("Go is fun"))
}
