package main

import (
	"context"
	"fmt"
	"sync"
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
				time.Sleep(1 * time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(i)
			}
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	wg.Wait()
	fmt.Println("Is it done?")
	time.Sleep(2 * time.Second)
	fmt.Println("Yes, its done")
}
