package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println(sum)

	m1 := map[string]int{"price": 40, "phone": 3}

	for key, value := range m1 {
		fmt.Println(key, value)
	}

	for key := range m1 {
		fmt.Println(key)
	}

	// give unicode code
	for i, c := range "GoLang" {
		fmt.Println(i, c, string(c))
	}

}
