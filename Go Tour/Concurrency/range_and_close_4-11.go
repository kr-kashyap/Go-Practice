package main

import (
	"fmt"
)

func print_i(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

func range_and_close_4_11() {
	c := make(chan int, 10)
	go print_i(cap(c), c)
	// This loop needs a call to close()
	for i := range c {
		fmt.Println(i)
	}

	/*
		However, in this case we don't need the close() call.
		for i := 0; i < cap(c); i++ {
			fmt.Println(i)
		}
	*/
}

/*
	? Needs some more exploration ?

	Whenever we are using a range we might need to make a call to the close() function. In here we can see
	we have definitive size of the buffered channel, and we are using range, which should ideally iterate
	through the capacity of the channel. And our loop in the function loop is also definitive. Still
	whenever we are omitting this close() call it panics.

	But when we use
		fori := 0 ; i < 10 ; i ++{} //We can omit the close() call.
*/
