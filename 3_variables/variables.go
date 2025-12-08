package main

import (
	"fmt"
)

func main() {
	// var name string = "Param"
	// var name = "Param"
	// var isAdult = true
	// var isAdult bool = true
	// var age = 24
	// var age int = 24

	// short hand syntax
	// name := "param"

	// by this you can declare variable and if you dont use that still that will not consider as error

	var _ string
	_ = "name 4"

	var _ string
	_ = "name 3"

	// fmt.Println(_) this will not allowed

	var name string
	name = "param"

	var age int
	age = 32

	var price float32
	price = 35.98

	fmt.Println(price)
	fmt.Println(age)
	fmt.Println(name)

	// fmt.Println("name " + name + " and is he adult : ")
}
