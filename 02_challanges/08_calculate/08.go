package main

import "fmt"

func main() {

	var (
		a, b int
		op   string
		res  float64
	)

	fmt.Println("enter value A")
	fmt.Scanln(&a)

	fmt.Println("enter value B")
	fmt.Scanln(&b)

	fmt.Println("enter operation (*, /, -, +):")
	fmt.Scanln(&op)

	// Map of functions
	m := map[string]func(int, int) float64{}

	m["*"] = multiplication
	m["+"] = addition
	m["-"] = subtract
	m["/"] = division

	// get function based on op
	fn, exists := m[op]

	if !exists {
		fmt.Println("Invalid operation")
		return
	}

	// call function
	res = fn(a, b)

	fmt.Println("Result:", res)
}

func addition(a, b int) float64 {
	return float64(a + b)
}

func subtract(a, b int) float64 {
	return float64(a - b)
}

func multiplication(a, b int) float64 {
	return float64(a * b)
}

func division(a, b int) float64 {
	if b == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	return float64(a) / float64(b)
}
