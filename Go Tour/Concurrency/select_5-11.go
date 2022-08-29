package main

import (
	"fmt"
	"time"
)

func print_i0(c chan int) {
	for i := 0; i <= cap(c); i++ {
		select {
		case c <- i:
			i = i
		}
	}
}

func select_5_11() {
	fmt.Println("\n--> Select 1 <--\n")
	select_5_11_1()
	c := make(chan int, 25)
	d := make(chan int, 25)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("C: %d\n", <-c)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("D: %d\n", <-d)
		}
	}()
	print_i0(c)
	print_i0(d)
}

/*
	Use of an Anonymous function in Goroutines.

	The select statement lets a goroutine wait on multiple communication operations.

	? A select blocks until one of its cases can run, then it executes that case. It chooses one at random if
	multiple are ready. ?

*/
