package main

import "fmt"

func main() {

	oddCount := 0
	evenCount := 0

	nums := []int{5, 2, 8, 7, 3, 10}

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evenCount = evenCount + 1
		} else {
			oddCount = oddCount + 1
		}
	}

	fmt.Println("Odd number = ", oddCount)
	fmt.Println("even number = ", evenCount)
}
