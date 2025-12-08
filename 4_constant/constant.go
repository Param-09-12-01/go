package main

import "fmt"

const age = 32 // allowed

// age2 := 2 not allowed

// like this we can also create the const
const (
	PORT = 3306
	url  = "Localhost"
)

func main() {

	// we can use const as variable
	// and if we dont use that will
	// not give the error

	const name string = "param"
	const name2 = "param"

	// if we declare outside method
	// and same we declare in method give erorr
	// var age = 32

	fmt.Println(PORT)
	fmt.Println(url)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name2)
}
