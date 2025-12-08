package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}

	fmt.Println(checkIfExists(arr, 01))
	fmt.Println(checkIfExists(arr, 10))
}

func checkIfExists(arr []int, target int) string {
	for _, val := range arr {
		if val == target {
			return "found"
		}
	}

	return "not found"
}
