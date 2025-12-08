package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go is very fast and simple"

	slen := strings.Split(s, " ")

	fmt.Println(len(slen))
}
