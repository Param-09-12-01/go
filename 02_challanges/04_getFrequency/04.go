package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 1}

	m := make(map[int]int)

	for _, val := range nums {

		value, ok := m[val]

		if ok {
			value = value + 1
			m[val] = value
		} else {
			m[val] = 1
		}
	}

	fmt.Println(m)

}
