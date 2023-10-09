package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/* Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться. */

func main() {
	c := make(chan int)
	N := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), N)
	defer cancel()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Publisher closed")
				close(c)
				return
			default:
				c <- i
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			num, ok := <-c
			if !ok {
				fmt.Println("Consumer closed")
				return
			}
			fmt.Printf("Got new mesage:%d\n", num)
		}
	}()

	wg.Wait()
}
