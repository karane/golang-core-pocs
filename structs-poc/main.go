package main

import (
	"fmt"
)

// 1. Define a struct
type Person struct {
	Name string
	Age  int
}

// 2. Struct methods
// Method with value receiver
func (p Person) Greet() {
	fmt.Printf("Greet: Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Method with pointer receiver
func (p *Person) HaveBirthday() {
	p.Age += 1
	fmt.Printf("HaveBirthday: Happy birthday, %s! You are now %d years old.\n", p.Name, p.Age)
}

// 4. Struct embedding (composition)
type Employee struct {
	Person     // embedding Person struct
	EmployeeID int
	Department string
}

func main() {
	// Example 1: Creating a struct instance
	fmt.Println("1. Define struct example")
	p1 := Person{Name: "Alice", Age: 30}
	p1.Greet()
	p1.HaveBirthday()
	p1.Greet()

	fmt.Println("\n\n2. Anonymous struct example")
	// Example 2: Anonymous struct
	project := struct {
		Title  string
		Budget int
	}{
		Title:  "Apollo",
		Budget: 100000,
	}
	fmt.Printf("Anonymous struct: Project: %s, Budget: %d\n", project.Title, project.Budget)

	fmt.Println("\n\n3. Struct embedding example")
	// Example 3: Using struct embedding
	e1 := Employee{
		Person:     Person{Name: "Bob", Age: 40},
		EmployeeID: 101,
		Department: "Engineering",
	}
	fmt.Printf("Employee: %s, Age: %d, ID: %d, Dept: %s\n",
		e1.Name, e1.Age, e1.EmployeeID, e1.Department)
	e1.Greet()
	e1.HaveBirthday()
	e1.Greet()

	fmt.Println("\n\n4. Struct embedding with slice example")
	// Example 4: Slice of embedded structs
	employees := []Employee{
		{Person: Person{Name: "Alice", Age: 30}, EmployeeID: 101, Department: "Engineering"},
		{Person: Person{Name: "Bob", Age: 40}, EmployeeID: 102, Department: "HR"},
		{Person: Person{Name: "Charlie", Age: 35}, EmployeeID: 103, Department: "Sales"},
	}

	for i, e := range employees {
		fmt.Printf("\n%d Employee: %s, Age: %d, ID: %d, Dept: %s\n", i+1, e.Name, e.Age, e.EmployeeID, e.Department)
		e.Greet()
		e.HaveBirthday()
	}
}
