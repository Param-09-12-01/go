package main

import "fmt"

func main() {

	marks := []int{75, 80, 65, 90, 55}
	totalMarks := 0

	for i := 0; i < len(marks); i++ {
		totalMarks = totalMarks + marks[i]
	}

	fmt.Println("total Marks = ", totalMarks)

	per := totalMarks / len(marks)

	fmt.Println("Average = ", per)

	if per >= 90 {
		fmt.Println("Grade = A")
	} else if per >= 75 {
		fmt.Println("Grade = B")
	} else if per >= 60 {
		fmt.Println("Grade = C")
	} else {
		fmt.Println("Grade = D")
	}
}
