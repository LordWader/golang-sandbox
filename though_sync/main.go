package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SharedMemory struct {
	m         *sync.RWMutex
	buffer    []int
	db        []int // simulate slice as db
	isBlocked atomic.Bool
	ch        chan int
}

func NewSharedMemory() *SharedMemory {
	return &SharedMemory{
		m:         &sync.RWMutex{},
		buffer:    make([]int, 0, 10),
		db:        make([]int, 0, 10000),
		isBlocked: atomic.Bool{},
		ch:        make(chan int, 1),
	}
}

func (nsm *SharedMemory) SimulateTraffic(initNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	init := initNumber
loop:
	for {
		select {
		case <-time.Tick(time.Millisecond * 200):
			nsm.ch <- init
			fmt.Printf("Normal traffic value: %d\n", init)
			init++
			if init == 200 {
				break loop
			}
		}
	}
	close(nsm.ch)
}

func (nsm *SharedMemory) MainEventProcessor(wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range nsm.ch {
		if blocked := nsm.isBlocked.Load(); blocked {
			fmt.Println("can't process normal traffic. Trying to put it to buffer")
			nsm.buffer = append(nsm.buffer, val)
			if len(nsm.buffer) >= 10 {
				if ok := nsm.m.TryLock(); ok {
					nsm.db = append(nsm.db, nsm.buffer...)
					nsm.buffer = nsm.buffer[:0]
					fmt.Println("buffer now empty")
					nsm.m.Unlock()
				}
			}
		} else {
			fmt.Println("processing as normal")
			// first we need to check remainings in buffer!
			if len(nsm.buffer) > 0 {
				fmt.Println("processing remainings from buffer")
				for _, val := range nsm.buffer {
					fmt.Printf("processed value from buffer: %d\n", val)
				}
				nsm.buffer = nsm.buffer[:0]
			}
			fmt.Printf("processing normal event: %d\n", val)
			fmt.Println("processed!")
		}
	}
}

func (nsm *SharedMemory) BulkEventProcessor(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 30)
	fmt.Println("Time to process buffers")
	nsm.m.Lock()
	for _, val := range nsm.db {
		fmt.Printf("processed value from db: %d\n", val)
	}
	nsm.m.Unlock()
	nsm.isBlocked.Store(false)
}

func main() {
	shm := NewSharedMemory()
	shm.isBlocked.Store(true)
	var wg sync.WaitGroup
	wg.Add(3)
	go shm.SimulateTraffic(0, &wg)
	go shm.MainEventProcessor(&wg)
	go shm.BulkEventProcessor(&wg)
	wg.Wait()
}
