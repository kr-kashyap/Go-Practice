package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	temp := ""
	for j := 0; j < 3; j++ {
		temp += strconv.Itoa(int(ip[j])) + "."
	}
	temp += strconv.Itoa(int(ip[3]))
	return temp
}

func exercise_stringers_18_26() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%16v: %10v\n", name, ip)
	}
}

/*
	String to int
	// import "strconv"
	strconv.Itoa(i) // i int
*/
