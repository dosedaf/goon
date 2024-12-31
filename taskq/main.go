package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	MaxWorker      int
	QueueSize      int
	QueuedTaskChan chan func()
	WaitGroup      sync.WaitGroup
	Ctx            context.Context
	Cancel         context.CancelFunc
}

func NewWorkerPool(queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		MaxWorker:      5,
		QueueSize:      queueSize,
		QueuedTaskChan: make(chan func(), queueSize),
		Ctx:            ctx,
		Cancel:         cancel,
	}
}

func (wp *WorkerPool) Start() {
}

func (wp *WorkerPool) AddTask(task func()) {
}

func main() {
	wp := NewWorkerPool()

	wp.Start()

	for start := time.Now(); time.Since(start) < 5*time.Second; {
		wp.AddTask(func() {
			fmt.Println("working on a task")
			time.Sleep(2 * time.Second)
			fmt.Println("done")
		})
	}
}
