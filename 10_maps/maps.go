package main

import (
	"fmt"
	"maps"
)

func main() {

	// var maps map[string]int
	// maps = make(map[string]int)
	// maps["goLang"] = 1
	// fmt.Println(maps["goLang"])
	// fmt.Println(maps["java"]) // if there is no key present then it will retuen default entry

	maps2 := make(map[string]string)

	maps2["GoLang"] = "backend"
	maps2["c++"] = "backend"
	fmt.Println(maps2["java"])
	fmt.Println(maps2["GoLang"])

	fmt.Println(len(maps2)) // to get length

	// delete one element
	delete(maps2, "java")
	fmt.Println(len(maps2))

	fmt.Println(maps2)

	// clear all element
	clear(maps2)

	fmt.Println(maps2)

	m := map[string]int{"price": 40, "phone": 3}
	fmt.Println(m)

	// to check element exits

	value, ok := m["price"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("price not present")
	}

	// how to check map is same or not

	m1 := map[string]int{"price": 40, "phone": 3}
	m2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(maps.Equal(m1, m2))

	// to iteate over map
	for key, value := range m {
		fmt.Println(key, value)
	}

}
