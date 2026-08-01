package main

import "testing"

func TestReverse(t *testing.T) {
	got := Reverse("golang")
	want := "gnalog"
	if got != want {
		t.Errorf("Reverse() = %q, want %q", got, want)
	}
}

func TestCountWords(t *testing.T) {
	got := CountWords("Go is fun")
	want := 3
	if got != want {
		t.Errorf("CountWords() = %d, want %d", got, want)
	}
}

// 2️Table-driven tests with SUBTESTS

func TestReverseSubtests(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "abc", "cba"},
		{"two letters", "Go", "oG"},
		{"empty string", "", ""},
		{"emoji", "👋🙂", "🙂👋"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountWordsSubtests(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"normal", "Go is fun", 3},
		{"multiple spaces", "  multiple   spaces  ", 2},
		{"empty", "", 0},
		{"single", "one", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.input)
			if got != tt.want {
				t.Errorf("CountWords(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// 3️Benchmarks (same as before)

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWords("Go is a statically typed compiled language")
	}
}
