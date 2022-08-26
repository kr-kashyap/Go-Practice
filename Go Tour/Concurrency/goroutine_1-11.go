package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
	Why does the second say function does not have the go keyword? Nd what
	if it does?

	NOTE: Here go say is the child thread. So the world may appear at max
	5 times. As it is the child thread, it can be terminated when the main
	thread is terminated.

	Output [Output will keep on varyin coz of context switch]

	world
	hello
	hello
	world
	hello
	world
	hello
	world
	world
	hello

*/
