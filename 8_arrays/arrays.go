package main

import "fmt"

// normally there is no array will use in go
//when to use array
//---------fixed size
//---------memory optimization
//---------constant time access
// number sequence of specific length
func main() {

	// for int default value 0
	// for bool default value false
	// for string that will be empty
	var nums [4]int
	fmt.Println(" length of nums ", len(nums))

	nums[3] = 4

	fmt.Println("last index of array is ", nums[len(nums)-1])
	fmt.Println("total array length value is ", nums)

	vals := [...]string{"param", "is", "learning", "go"}
	fmt.Println(vals)

	vals2 := [len(vals)]int{1, 3, 3, 4}
	fmt.Println(vals2)
	vals2[0] = 4
	fmt.Println(vals2)

	twoDArray := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDArray)
}
