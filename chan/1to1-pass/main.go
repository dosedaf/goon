package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("working...")
		done <- struct{}{}
	}()

	<-done // blocks until there's something to pass, line 14
	fmt.Println("done")
}
