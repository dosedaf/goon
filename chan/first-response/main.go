package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func source(ch chan<- int32) {
	time.Sleep(2 * time.Second)
	ch <- rand.Int32N(100)
}

func main() {
	startTime := time.Now()

	ch := make(chan int32, 5)

	for i := 0; i < cap(ch); i++ {
		go source(ch)
	}

	rnd := <-ch

	fmt.Println(time.Since(startTime))

	fmt.Println(rnd)
}
