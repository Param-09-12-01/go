package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 5}

	m := make(map[int]bool)

	unique := make([]int, 0)

	for _, vals := range nums {
		_, ok := m[vals]

		if !ok {
			m[vals] = true
			unique = append(unique, vals)
		}
	}

	fmt.Println(unique)
}
