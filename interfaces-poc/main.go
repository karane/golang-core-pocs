package main

import (
	"fmt"
)

// 1. Define the first interface
type Speaker interface {
	Speak() string
}

// 2. Define a second interface
type Walker interface {
	Walk() string
}

// 3. Define a struct that implements both interfaces implicitly
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

func (d Dog) Walk() string {
	return d.Name + " is walking happily."
}

// 4. Define another struct that only implements Speaker
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow! I'm " + c.Name
}

func main() {
	// 5. Using the Speaker interface
	var s Speaker
	s = Dog{Name: "Rex"}
	fmt.Println("5. Dog as Speaker ->", s.Speak())

	s = Cat{Name: "Mittens"}
	fmt.Println("5. Cat as Speaker ->", s.Speak())

	// 6. Using the Walker interface
	var w Walker = Dog{Name: "Rex"}
	fmt.Println("6. Dog as Walker ->", w.Walk())

	// 7. Dog can be assigned to both Speaker and Walker
	d := Dog{Name: "Buddy"}
	fmt.Println("7. Dog as Speaker ->", d.Speak())
	fmt.Println("7. Dog as Walker ->", d.Walk())

	// 8. Cat only implements Speaker
	var s2 Speaker = Cat{Name: "Mittens"}
	fmt.Println("8. Cat as Speaker ->", s2.Speak())
	// var w2 Walker = Cat{Name: "Mittens"} // compile error

	// 9. Type assertion: checking concrete type from an interface
	if cat, ok := s2.(Cat); ok {
		fmt.Println("9. Type assertion succeeded -> Cat name is:", cat.Name)
	}

	// 10. Empty interface can hold any type
	var any interface{}
	any = 42
	fmt.Println("10. any holds an int:", any)

	any = "Hello Go!"
	fmt.Println("10. any now holds a string:", any)

	// 11. Type switch with empty interface
	switch v := any.(type) {
	case int:
		fmt.Println("11. It's an int with value:", v)
	case string:
		fmt.Println("11. It's a string with value:", v)
	default:
		fmt.Println("11. Unknown type")
	}

	// 12. Multiple interface type assertion
	var i interface{} = Dog{Name: "Charlie"}
	if sp, ok := i.(Speaker); ok {
		fmt.Println("12. Dog asserted as Speaker ->", sp.Speak())
	}
	if wa, ok := i.(Walker); ok {
		fmt.Println("12. Dog asserted as Walker ->", wa.Walk())
	}

	// 13. Interface embedding (combine multiple interfaces)
	type Animal interface {
		Speaker
		Walker
	}

	var a Animal = Dog{Name: "Max"}
	fmt.Println("13. Animal as Speaker ->", a.Speak())
	fmt.Println("13. Animal as Walker ->", a.Walk())
}
