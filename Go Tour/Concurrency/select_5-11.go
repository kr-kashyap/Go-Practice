package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func select_5_11() {
	c := make(chan int)
	d := make(chan int)
	quit := make(chan int)
	//quit1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("D: %d\n", <-d)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	fibonacci(d, quit)
}
