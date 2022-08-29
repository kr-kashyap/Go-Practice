package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := -1, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func exercise_26_27() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fib(%2d) = %3d\n", i, f())
	}
}
