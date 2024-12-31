package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(3 * time.Second)
		r <- rand.Int32N(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	a, b := longTimeRequest(), longTimeRequest()

	fmt.Println(sumSquares(<-a, <-b))
}
