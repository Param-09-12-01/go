package main

import (
	"fmt"
	"slices"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}
	slices.Reverse(arr2)
	fmt.Println(arr2)
	fmt.Println(reverse(arr1))
}

func reverse(arr []int) []int {
	arrLength := len(arr)
	for i := 0; i < arrLength/2; i++ {
		swap(i, arrLength-1-i, arr)
		fmt.Println(arr)
	}
	return arr
}

func swap(fromIndex int, toIndex int, arr []int) {
	temp := arr[toIndex]
	arr[toIndex] = arr[fromIndex]
	arr[fromIndex] = temp
}
