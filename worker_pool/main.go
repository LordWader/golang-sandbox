package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const POOL_SIZE = 10

type Worker struct {
	id      int
	in      chan int
	ctx     context.Context
	handler func(w *Worker, val int)
}

type WorkerPool struct {
	workers []*Worker
	wg      sync.WaitGroup
}

func NewWorkerPool(arr []*Worker) *WorkerPool {
	return &WorkerPool{
		workers: arr,
	}
}

func (wp *WorkerPool) Run() {
	for _, worker := range wp.workers {
		wp.wg.Add(1)
		go func(w *Worker) {
			w.Process(&wp.wg)
		}(worker)
	}
	wp.wg.Wait()
}

func NewWorker(in chan int, ctx context.Context, id int, handler func(w *Worker, val int)) *Worker {
	return &Worker{
		in:      in,
		ctx:     ctx,
		id:      id,
		handler: handler,
	}
}

func (w *Worker) Process(wg *sync.WaitGroup) {
	fmt.Printf("Worker %d started procesing numbers\n", w.id)
	for {
		select {
		case <-w.ctx.Done():
			wg.Done()
			fmt.Printf("Worker %d, done processing\n", w.id)
			return
		case val, ok := <-w.in:
			if !ok {
				wg.Done()
				fmt.Printf("Channel close! Worker %d stopped!\n", w.id)
				return
			}
			w.handler(w, val)
		}
	}
}

func Task(w *Worker, val int) {
	time.Sleep(time.Second * 3)
	fmt.Printf("Worker %d, processed num: %d\n", w.id, val)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int, POOL_SIZE)
	workers := make([]*Worker, 0, POOL_SIZE)
	for i := 0; i < POOL_SIZE; i++ {
		workers = append(workers, NewWorker(in, ctx, i, Task))
	}
	workerPool := NewWorkerPool(workers)
	go func() {
		defer close(in)
		i := 0
	loop:
		for i < 500 {
			select {
			case <-ctx.Done():
				break loop
			case in <- i:
				i++
			case <-time.Tick(time.Second * 2):
				fmt.Println("Overflow!")
				cancel()
			}
		}
	}()
	//simulate shutdown
	go func() {
		min := 10
		max := 30
		for {
			r := rand.Intn(max-min+1) + min
			if r%10 == 0 {
				fmt.Println("Sending done signal")
				cancel()
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	workerPool.Run()
}
