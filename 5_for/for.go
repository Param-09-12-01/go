package main

import (
	"fmt"
)

// for -> only loop in go
func main() {

	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// normal loop

	// for i := 0; i < 4; i++ {

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)

	// 	if i == 1 {
	// 		break
	// 	}
	// }

	// to print 1 to 9
	for i := range 10 {
		fmt.Println(i)
	}

}
