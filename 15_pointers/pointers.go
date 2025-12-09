package main

import "fmt"

func main() {
	fmt.Println("=== Pass-by-value demo ===")
	val := 1
	printAddrAndValue("main.val (before)", &val)
	demoByValue(val)
	printAddrAndValue("main.val (after)", &val)

	fmt.Println()
	fmt.Println("=== Pass-by-pointer demo ===")
	ptrVal := 10
	printAddrAndValue("main.ptrVal (before)", &ptrVal)
	demoByPointer(&ptrVal)
	printAddrAndValue("main.ptrVal (after)", &ptrVal)
}

// demoByValue receives a copy of the integer.
// Modifying the parameter does NOT change the caller's variable.
func demoByValue(n int) {
	printAddrAndValue("demoByValue.n (entry)", &n)
	n = 5
	fmt.Printf("demoByValue: changed local n to %d\n", n)
	printAddrAndValue("demoByValue.n (exit)", &n)
}

// demoByPointer receives a pointer to an int.
// Modifying *n changes the caller's variable.
func demoByPointer(n *int) {
	// print pointer itself (%p) and the value it points to
	fmt.Printf("demoByPointer: received pointer %p -> value %d\n", n, *n)
	*n = 5
	fmt.Printf("demoByPointer: changed *n to %d\n", *n)
}

// helper to print the address and the value at that address
func printAddrAndValue(label string, p *int) {
	fmt.Printf("%-30s addr=%p  value=%d\n", label, p, *p)
}
