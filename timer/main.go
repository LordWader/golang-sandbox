package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func heavyCalc() int {
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	return 42
}

func process(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan int, 1)

	go func() {
		defer close(ch)
		select {
		case <-ctx.Done():
			fmt.Println("Done by context")
		default:
			ch <- heavyCalc()
		}
	}()
	select {
	case res := <-ch:
		fmt.Printf("Got result from heavy calc: %d\n", res)
	case <-time.After(time.Second * 2):
		fmt.Println("Timed out!")
	}
}

func main() {
	ctx := context.Background()
	i := 0
	for i < 10 {
		process(ctx)
		i++
	}
}
