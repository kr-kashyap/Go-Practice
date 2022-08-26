package main

import "fmt"

func main() {
	fmt.Println("\n--> GoRoutine <--\n")
	//go goroutine_1_11()
	goroutine_1_11()
	fmt.Println("\n--> Channels <--\n")
	channels_2_11()
	fmt.Println("\n--> Buffered Channels <--\n")
	buffered_channels()
	fmt.Println("\n--> Range and Close <--\n")
	range_and_close_4_11()
	fmt.Println("\n--> Select <--\n")
	select_5_11()
	fmt.Println("\n--> Default Select <--\n")
	default_selection_6_11()
	fmt.Println("\n--> SyncMutex <--\n")
	syncMutex_9_11()
}
