package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func consumeBatch(ctx context.Context, in <-chan int, stop <-chan bool) {
	defer wg.Done()
	for {
		select {
		case _ = <-stop:
			fmt.Printf("Length of channel: %d \n", len(in))
			for len(in) > 0 {
				fmt.Printf("value from channel after finish: %d \n", <-in)
			}
			return
		case _ = <- ctx.Done():
			fmt.Printf("Length of channel: %d\n", len(in))
			for len(in) > 0 {
				fmt.Printf("value from channel after cancelling: %d \n", <-in)
			}
			return
		default:
			if len(in) < 9 {
				time.Sleep(time.Second)
				fmt.Printf("sleep and get len channel: %d \n", len(in))
			} else {
				for i := 0; i < 10; i++ {
					fmt.Printf("value from channel: %d \n", <-in)
				}
			}
		}
	}
}

func produceBatch(ctx context.Context, in chan<- int, stop chan<- bool) {
	var i int
	defer wg.Done()
	for i < 30 {
		select {
		case _ = <-ctx.Done():
			//stop <- true
			fmt.Println("Interrupting, closing goroutine")
			return
		default:
			in <- i
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Adding num %d to channel\n", i)
			i++
		}
	}
	stop <- true
	fmt.Println("Finish producing messages")
	return
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	wg.Add(2)
	input := make(chan int, 15)
	stop := make(chan bool)
	go produceBatch(ctx, input, stop)
	go consumeBatch(ctx, input, stop)
	wg.Wait()
	fmt.Printf("Done with batch processing of events")
}
