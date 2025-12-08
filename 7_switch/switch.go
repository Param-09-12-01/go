package main

import (
	"fmt"
	"time"
)

func main() {

	i := 12

	switch i {
	case 1:
		fmt.Println("one", i)
	case 2:
		fmt.Println("two", i)
	default:
		fmt.Println("other", i)
	}

	// multiple condition

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend", time.Now().Weekday())
	default:
		fmt.Println("weekday", time.Now().Weekday())
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch varType := i.(type) {
		case int:
			fmt.Println("this is int ", varType)
		case bool:
			fmt.Println("this is bool ", varType)
		case string:
			fmt.Println("this is string ", varType)
		default:
			fmt.Println("other type ", varType)
		}
	}

	// whoAmI("param")
	// whoAmI(true)
	whoAmI(45)

}
