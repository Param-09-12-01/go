package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
// most use construct in go
// have some useful methods
func main() {

	// default uniitilized slice is nil
	var nums []int

	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	// how to make non nil slice
	var numsNonNull = make([]int, 2)
	fmt.Println(cap(numsNonNull))
	fmt.Println(numsNonNull)
	// this give error because when you append then
	//only go expand the size of array
	// numsNonNull[4] = 6
	fmt.Println(numsNonNull)

	// first parameter is type , second is length , and third is capacity
	var numsNonNullWithDiffCapacity = make([]int, 2, 5)
	// this will also give errir cause len = 2 and capacity = 5
	// so index out of range
	// numsNonNullWithDiffCapacity[4] = 4
	fmt.Println(numsNonNullWithDiffCapacity)

	// after reaching the capacity the capacity will double
	var numsWithAppend = make([]int, 2, 5)
	numsWithAppend = append(numsWithAppend, 2)

	fmt.Println("capacity of numsWithAppend is ", cap(numsWithAppend))
	numsWithAppend = append(numsWithAppend, 2)
	numsWithAppend = append(numsWithAppend, 3)
	numsWithAppend = append(numsWithAppend, 4)
	numsWithAppend = append(numsWithAppend, 5)

	fmt.Println("capacity of numsWithAppend after appending char is ", cap(numsWithAppend))

	fmt.Println(numsWithAppend) // this will append at 2nd index

	// to create short hand slice
	numsWithShortHand := []int{}
	fmt.Println(numsWithShortHand)

	// to insert with index
	var numsInsertWithIndex = make([]int, 2, 5)
	numsInsertWithIndex[0] = 4
	fmt.Println(numsInsertWithIndex)

	// declare and assign witout making
	// var numsToCopy []int
	// numsToCopy[1] = 4
	// fmt.Println(numsToCopy)

	// Create a slice with len=4, cap=10
	numsToCopy := make([]int, 4, 10)
	// Create copy slice with length = current length (4)
	numsWhereWeCreateCopy := make([]int, len(numsToCopy))
	// Append increases LENGTH (now len=6). Capacity stays 10.
	numsToCopy = append(numsToCopy, 3, 4)
	fmt.Println(numsToCopy) // [0 0 0 0 3 4]
	// copy(dest, src) copies min(len(dest), len(src))
	// dest len = 4, src len = 6 → only first 4 elements copied
	copy(numsWhereWeCreateCopy, numsToCopy)
	fmt.Println("original array : ", numsToCopy)          // [0 0 0 0 3 4]
	fmt.Println("copied array : ", numsWhereWeCreateCopy) // [0 0 0 0]

	// code where actually copies
	numsToCopyOG := make([]int, 4, 10)
	numsToCopyOG = append(numsToCopyOG, 3, 4)
	// Create destination AFTER append, so len=6
	numsWhereWeCreateCopyOG := make([]int, len(numsToCopyOG))
	copy(numsWhereWeCreateCopyOG, numsToCopyOG)
	fmt.Println("original array:", numsToCopyOG)          // [0 0 0 0 3 4]
	fmt.Println("copied array:", numsWhereWeCreateCopyOG) // [0 0 0 0 3 4]

	//slice copy
	var sliceExe = []int{1, 2, 3, 4}
	fmt.Println(sliceExe[0:2]) // give op till 1st index Give the before index
	fmt.Println(sliceExe[:2])  // give till 1st index
	fmt.Println(sliceExe[2:])  // starts from  2nd index
	// fmt.Println(sliceExe[2:6]) // give error
	fmt.Println(sliceExe[0:4]) // give op

	nums1st := []int{1, 2}
	nums2st := []int{1, 2}
	fmt.Println(slices.Equal(nums1st, nums2st))

	var (
		nums3 []int
		nums4 []int
	)

	fmt.Println(slices.Equal(nums3, nums4))

	// to create 2D slices
	nums2D := [][]int{{1, 2, 3}, {2, 2, 3, 4}}
	fmt.Println(nums2D)
}
