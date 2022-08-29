package main

import (
	"fmt"
	"time"
)

func fibonacci_1(c, quit chan int) {
	x := 0
	for {
		select {
		case c <- x:
			x++
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func select_5_11_1() {
	fmt.Println("\n NOTE : In here we have Empty channels. I.E. They are not initialized for a specific length. And also in the function we have")
	fmt.Println("used a Forever loop. So we need to make a breaking condition by passing value to the quit channel.\n")
	// Size of the channel is not initialized here.
	c1 := make(chan int)
	// Size of the channel is not initialized here.
	d1 := make(chan int)
	quit := make(chan int)
	//quit1 := make(chan int)
	go func() {
		for i := 0; i < 25; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("C1: %d\n", <-c1)
		}
		quit <- 0
	}()
	go func() {
		for i := 0; i < 25; i++ {
			fmt.Printf("D1: %d\n", <-d1)
		}
		quit <- 0
	}()
	fibonacci_1(c1, quit)
	fibonacci_1(d1, quit)
}

/*
	Use of an Anonymous function in Goroutines.

	The select statement lets a goroutine wait on multiple communication operations.

	NOTE : In here we have Empty channels. I.E. They are not initialized for a specific length. And also in the function we have
	used a Forever loop. So we need to make a breaking condition by passing value to the quit channel.

	? A select blocks until one of its cases can run, then it executes that case. It chooses one at random if
	multiple are ready. ?

*/
