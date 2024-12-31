package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		fmt.Println("working")
		time.Sleep(1 * time.Second)

		<-done // receive value from channel
	}()

	// buffered channels need receiver, it waits for the goroutine, in case it receives value from the channel
	done <- struct{}{}
	fmt.Println("done")
}
