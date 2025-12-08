package main

import "fmt"

// go does not have ternary operator we need to use regular if else in Go
func main() {

	age := 12

	if age == 18 {
		fmt.Println("Person is adult 18 year old ")
	} else if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	fmt.Println("complete if else")

	fmt.Println("====================            logical opeartor            ====================")

	role := "Admin"

	if role == "Admin" && false {
		fmt.Print("is Admin true")
	}

	// declaring variable in if condition
	if age := 15; age >= 18 {
		fmt.Print("person is an adult ", age)
	} else if age >= 12 {
		fmt.Print("person is teenager ", age)
	} else {
		fmt.Print("in else condition ", age)
	}
}
