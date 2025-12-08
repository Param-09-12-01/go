package main

import "fmt"

func main() {

	fmt.Println(multiply(4, 5))
}

func multiply(num int, timesMultiply int) int {

	res := 1
	for range timesMultiply {
		res = res * num
	}
	fmt.Println("Go version supports range-int loops ✔")

	return res
}
