package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func consumeBatch(ctx context.Context, in <-chan int) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done) // no need to send anything here, close is an event itself

		buf := make([]int, 0, 10)

		defer func() {
			fmt.Println("Done consuming")
			fmt.Printf("Length of buffer: %d; channel: %d\n", len(buf), len(in))
			for _, v := range buf {
				fmt.Printf("value from buffer: %d \n", v)
			}
			for len(in) > 0 {
				fmt.Printf("value from channel: %d \n", <-in)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				buf = append(buf, v)

				if len(buf) < 10 {
					time.Sleep(time.Second)
					fmt.Printf("sleep and get len buffer: %d \n", len(buf))
				} else {
					for i, v := range buf {
						fmt.Printf("value %d from channel: %d \n", i, v)
					}
					buf = buf[:0]
				}
			}
		}
	}()

	return done
}

func produceBatch(ctx context.Context) <-chan int {
	out := make(chan int) // make it unbuffered so no events are lost on cancel

	go func() {
		defer close(out)

		for i := 0; i < 30; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Interrupting, closing goroutine")
				return
			default:
				out <- i
				time.Sleep(time.Millisecond * 500)
				fmt.Printf("Adding num %d to channel\n", i)
			}
		}
		fmt.Println("Finish producing messages")
	}()

	return out
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	input := produceBatch(ctx)
	done := consumeBatch(ctx, input)

	<-done

	fmt.Printf("Done with batch processing of events")
}
