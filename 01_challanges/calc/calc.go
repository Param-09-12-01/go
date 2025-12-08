package main

import "fmt"

func main() {

	var i, j int
	var opration string

	fmt.Println("enter first number : ")
	fmt.Scan(&i)

	fmt.Println("enter second number : ")
	fmt.Scan(&j)

	fmt.Println("enter any arithmatic opration that you want to perform ( * , + , - , / , %) ")
	fmt.Scan(&opration)

	switch opration {
	case "*":
		fmt.Println("output of multiplication(*) is : ", i*j)
	case "+":
		fmt.Println("output of addition(+) is : ", i+j)
	case "-":
		fmt.Println("output of subtraction(-) is : ", i-j)
	case "/":
		if j == 0 {
			fmt.Println("can not devide by 0")
			break
		}
		fmt.Println("output of division(/) is : ", i/j)
	case "%":
		fmt.Println("output of percentage(*) is : ", (i*j)/100)
	default:
		fmt.Println("enter valid operation")
	}
}
