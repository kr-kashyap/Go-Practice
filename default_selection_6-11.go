package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
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
	O/P?
*/
