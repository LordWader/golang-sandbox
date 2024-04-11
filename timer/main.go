package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func heavyCalc(ctx context.Context, in chan int) {
	r := rand.Intn(3)
	select {
	case <-time.Tick(time.Duration(r)*time.Second + time.Millisecond*500):
		in <- 42
	case <-ctx.Done():
		fmt.Println("Exiting calculation goroutine!")
	}
}

func process(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan int, 1)

	go heavyCalc(ctx, ch)
	select {
	case res := <-ch:
		fmt.Printf("Got result from heavy calc: %d\n", res)
	case <-time.Tick(time.Second * 2):
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
