package main

import (
	"context"
	"fmt"
	"time"
)

/* Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться. */

func main() {
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	N := 5 * time.Second

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Publisher closed")
			default:
				c <- i
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Publisher closed")
			default:
				num := <-c
				fmt.Printf("Got new mesage:%d\n", num)
			}
		}
	}(ctx)

	time.Sleep(N)
	cancel()
}
