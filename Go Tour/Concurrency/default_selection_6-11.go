package main

import (
	"fmt"
	"time"
)

func default_selection_6_11() {
	fmt.Printf("%d\n", time.Millisecond)
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
	i := 1
	j := 1
	k := 1
	for {
		select {
		case <-tick:
			fmt.Printf("%2d tick.\n", j)
			j++
		case <-boom:
			fmt.Printf("%2d BOOM!\n", k)
			k++
			return
		default:
			fmt.Printf("%2d     .\n", i)
			i++
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*
	What is Tick and After?
	Tick repeats it at X intervals of time.
	After waits for x intervals of time.
	In here even if we use tick / after for Boom it does not matter coz we have return statement in the select of Boom.
	O/P?
*/
