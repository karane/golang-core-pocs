package greet

import "fmt"

// Exported function (starts with uppercase)
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Unexported function (starts with lowercase)
func whisperSecret(name string) {
	fmt.Printf("(whispering) %s has a secret...\n", name)
}

// Exported struct
type Person struct {
	Name string // Exported field
	age  int    // unexported field
}

// Exported constructor function
func NewPerson(name string, age int) Person {
	return Person{Name: name, age: age}
}

// Exported method (uppercase)
func (p Person) Show() {
	fmt.Printf("Person{name=%s, age=%d}\n", p.Name, p.age)
}

// Unexported method (lowercase)
func (p Person) secret() {
	fmt.Println("this is a secret method")
}
