package main

import (
	"log"
	"time"
)

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // wait for main to give value to ready

	log.Print("Worker#", id, " starts")
	time.Sleep(time.Second * 2)
	log.Print("Worker#", id, " done")

	done <- T{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)

	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	log.Print("program  starting")

	/*
		synced workers start
		the flow is,
		<- ready in worker will block until channel is closed / T{} is passed to the channel
		both allows <- ready to be done, thus continuing the code
	*/

	// this serves the purpose of syncing the workers to start after certain times, by blocking using ready
	time.Sleep(time.Second * 3 / 2)
	/*
		// this done sequentially, what if there are hunders of workers? won't the last worker start last?
			ready <- T{}
			ready <- T{}
			ready <- T{}
	*/
	// making use of closed channel feature in which you can receive infinitely
	close(ready)

	<-done // done blocking, wait for value from worker
	<-done
	<-done

	log.Print("done")
}
