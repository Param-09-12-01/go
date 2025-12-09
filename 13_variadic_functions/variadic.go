package main

import (
	"fmt"
)

func main() {
	res := sum(0, 20, 30, 40)
	fmt.Println("res1 ", res)
	res2 := sum(0, 20, 30, 40, 50, 60)
	fmt.Println("res2 ", res2)

	slices := []int{4, 5, 6, 7}
	res3 := sum(slices...)
	fmt.Println("res3 ", res3)

	test(10, "param", "is", "learning", "Go")

	test2(10, "param", 40.96)

}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total = total + val
	}
	return total
}

// we can not do this
// func test(nums ...int , str ...string)  {

// }

// this is valid
func test(nums int, str ...string) {
	fmt.Println(nums)
	fmt.Println(str)
}

// by this we can take any thing and as many as arg we want
func test2(anyType ...interface{}) {

	fmt.Println(anyType...)
	fmt.Println(len(anyType))

}
