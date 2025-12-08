package main

import "fmt"

func main() {

	arr := [6]int{12, 5, 7, 25, 1, 9}

	min := arr[0]
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}

		if min > arr[i] {
			min = arr[i]
		}
	}

	fmt.Println("max=", max)
	fmt.Println("min=", min)

}
