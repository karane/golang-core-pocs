package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // creates a basic error
	}
	return a / b, nil
}

// Define a custom error type for better context
type NotFoundError struct {
	Resource string
	ID       int
}

// IMPORTANT: Implement the `Error()` method to satisfy the `error` interface
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("resource '%s' with ID %d not found", e.Resource, e.ID)
}

// NotFoundErrorWithIs behaves like NotFoundError but implements Is(),
// so errors.Is can match it regardless of field values.
type NotFoundErrorWithIs struct {
	Resource string
	ID       int
}

func (e *NotFoundErrorWithIs) Error() string {
	return fmt.Sprintf("resource '%s' with ID %d not found", e.Resource, e.ID)
}

func (e *NotFoundErrorWithIs) Is(target error) bool {
	_, ok := target.(*NotFoundErrorWithIs)
	return ok
}

// Function that returns a custom error
func getResource(id int) (string, error) {
	if id != 1 {
		return "", &NotFoundError{Resource: "User", ID: id}
	}
	return "John Doe", nil
}

// Another layer of function that wraps lower-level errors
func findUser(id int) (string, error) {
	user, err := getResource(id)
	if err != nil {
		// Wrap error with additional context using %w (important!)
		return "", fmt.Errorf("findUser failed: %w", err)
	}
	return user, nil
}

func main() {
	fmt.Println("== Basic error ==")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("\n== Custom error ==")
	_, err = getResource(42)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n== Wrapping errors ==")
	// _, err = findUser(42)
	_, err = getResource(42)
	if err != nil {
		fmt.Println("Wrapped Error:", err)

		// errors.Is does a strict comparison (or calls Is() if implemented).
		// NotFoundError has no Is() method, so this compares &NotFoundError{}
		// against the actual error by identity/value equality and NEVER matches.
		if errors.Is(err, &NotFoundError{}) {
			fmt.Println("Detected NotFoundError using errors.Is")
		} else {
			fmt.Println("errors.Is did NOT detect NotFoundError (no Is() method, so it falls back to equality)")
		}

		// Use errors.As to extract the original custom error
		var notFound *NotFoundError
		if errors.As(err, &notFound) {
			fmt.Printf("Extracted custom error: Resource=%s, ID=%d\n", notFound.Resource, notFound.ID)
		}
	}

	fmt.Println("\n== errors.Is with a custom Is() method ==")
	err = &NotFoundErrorWithIs{Resource: "User", ID: 42}
	wrapped := fmt.Errorf("findUser failed: %w", err)
	if errors.Is(wrapped, &NotFoundErrorWithIs{}) {
		fmt.Println("Detected NotFoundErrorWithIs using errors.Is (thanks to its Is() method)")
	}
}
