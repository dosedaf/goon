package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func longTimeRequest(r chan<- int32) {
	time.Sleep(3 * time.Second)
	r <- rand.Int32N(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	// ca, cb := make(chan int32), make(chan int32)
	ch := make(chan int32)

	go longTimeRequest(ch)
	go longTimeRequest(ch)

	// go longTimeRequest(ca)
	// go longTimeRequest(cb)

	fmt.Println(sumSquares(<-ch, <-ch))
}
