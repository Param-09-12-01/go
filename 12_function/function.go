package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func getLanguage() (string, string, string) {
	return "Go", "Java", "c"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt1() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	res := add(10, 15)
	fmt.Println(res)

	resSub := sub(10, 15)
	fmt.Println(resSub)

	// resMultipleReturn :=
	fmt.Println(getLanguage())

	lang1, lang2, lang3 := getLanguage()
	fmt.Println(lang1, lang2, lang3)

	fmt.Println(lang1)

	f := func(a int) int {
		return a

	}
	processIt(f)

	processIt1()
}
