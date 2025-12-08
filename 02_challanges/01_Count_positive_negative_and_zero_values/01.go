package main

import "fmt"

func main() {
	nums := []int{2, -1, 0, 5, -7, 8, 0}
	zero, negetive, positive := giveCount(nums)

	fmt.Println("zero ", zero)
	fmt.Println("negetive ", negetive)
	fmt.Println("positive ", positive)
}

func giveCount(nums []int) (int, int, int) {

	zero, negetive, positive := 0, 0, 0

	for i := range nums {

		if nums[i] == 0 {
			zero = zero + 1
		} else if nums[i] > 0 {
			positive = positive + 1
		} else {
			negetive = negetive + 1
		}

	}

	return zero, negetive, positive
}
