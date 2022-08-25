package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	/*
		ch <- 3
		ch <- 4
	*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
	Prints only for the exact same length of the buffer. If it is
	lesser or greater it produces fatal error.
*/
