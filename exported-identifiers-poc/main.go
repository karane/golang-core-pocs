package main

import (
	"fmt"
	"myapp/greet"
)

func main() {
	greet.Hello("Karane") // OK: exported function

	// greet.whisperSecret("Karane")
	// ERROR: not exported

	p := greet.NewPerson("Maria", 30)
	p.Show() // OK: exported method

	fmt.Println(p.Name) // OK: exported field

	// fmt.Println(p.age)
	// ERROR: unexported field

	// p.secret()
	// ERROR: unexported method
}
