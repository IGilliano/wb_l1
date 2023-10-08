package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины

func main() {
	//Метод 1: Послать стоп-сигнал в канал
	quit1 := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			select {
			case <-quit1:
				return
			default:
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	quit1 <- true
	fmt.Println("Is it done?")
	time.Sleep(2 * time.Second)
	fmt.Println("Yes, its done")

	//Метод 2: Закрыть канал
	quit2 := make(chan string)
	go func() {
		for {
			num, ok := <-quit2
			if !ok {
				fmt.Println("Goroutine stopped")
				return
			}
			println(num)
		}
	}()
	quit2 <- "5"
	quit2 <- "4"
	quit2 <- "3"
	quit2 <- "2"
	quit2 <- "1"
	close(quit2)
	time.Sleep(5 * time.Second)

	//Метод 3: Использовать контекст
	quit3 := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				quit3 <- struct{}{}
				return
			default:
				fmt.Println(i)
			}
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("Is it done?")
	time.Sleep(2 * time.Second)
	fmt.Println("Yes, its done")
}
