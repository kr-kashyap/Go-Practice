package main

import "fmt"

func buffered_channels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//Trying to enter 3rd value.
	fmt.Printf("\nIn buffered channel ch with cap = %d. We are trying to insert after len = %d. Thus we get below message based on the if else check.\n", cap(ch), len(ch))
	if len(ch) < cap(ch) {
		ch <- 3
	} else {
		fmt.Println("Buffer channel is already full.")
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
