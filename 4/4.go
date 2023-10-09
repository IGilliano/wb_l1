package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.*/

func Tasks(data <-chan string, i int) {
	for {
		msg, ok := <-data
		if !ok {
			break
		}
		fmt.Fprintf(os.Stdout, msg, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(os.Stdout, "Goroutine is dead\n")
}

func main() {
	var workerPool int
	fmt.Println("Choose workers number")

	data := make(chan string, 100)
	exit := make(chan os.Signal, 1)
	defer close(exit)
	signal.Notify(exit, syscall.SIGINT)
	wg := sync.WaitGroup{}

	_, err := fmt.Scanf("%d", &workerPool)
	if err != nil {
		log.Fatalf("Incorrect input: %v", err.Error())
	}

	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			Tasks(data, i)
		}()
	}

	func() {
		for {
			select {
			case <-exit:
				close(data)
				return
			default:
				data <- "Hello, world!\n"
			}
		}
	}()

	wg.Wait()
	fmt.Println("Program is graceful shutdown")
}
