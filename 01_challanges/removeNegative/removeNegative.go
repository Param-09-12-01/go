package main

import "fmt"

func main() {

	nums := []int{5, -1, 7, -9, 3, -2}
	var nonNegtiveSlice []int

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			nonNegtiveSlice = append(nonNegtiveSlice, nums[i])
		}
	}

	fmt.Println(nonNegtiveSlice)

}
