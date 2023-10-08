package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.*/

func Tasks(data <-chan string, wg *sync.WaitGroup) {
	for range data {
		msg := <-data
		fmt.Fprintf(os.Stdout, msg)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(os.Stdout, "Goroutine is dead\n")
	wg.Done()
}

func main() {
	data := make(chan string)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT)
	wg := sync.WaitGroup{}

	workerPool := 3
	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		go Tasks(data, &wg)
	}

	func() {
		for {
			select {
			case <-exit:
				return
			default:
				data <- "Hello, world!\n"
			}
		}
	}()
	close(data)
	close(exit)
	wg.Wait()
	fmt.Println("Program is graceful shutdown")
}
