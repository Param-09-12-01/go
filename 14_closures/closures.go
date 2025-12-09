package main

import "fmt"

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	increment2, reset := counterWithReset()
	fmt.Println(increment2())
	fmt.Println(increment2())
	reset()
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
}

// counter returns a closure that "remembers" the value of count.
// Each call to the returned function increments and returns the
// updated value of count.
//
// This works because the inner function forms a closure over the
// 'count' variable. Even after counter() finishes executing,
// the variable 'count' remains alive in memory and is shared
// across all calls to the returned function, allowing the value
// to persist and update over time.
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

// counter returns two functions: an increment function and a reset function.
// The increment function forms a closure over 'count' and updates its value
// on each call. The reset function sets 'count' back to zero. Both functions
// share the same enclosed variable, so they operate on the same state.
func counterWithReset() (func() int, func()) {
	count := 0

	inc := func() int {
		count++
		return count
	}

	reset := func() {
		count = 0
	}

	return inc, reset
}
