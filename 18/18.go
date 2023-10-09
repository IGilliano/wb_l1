package main

import (
	"fmt"
	"sync"
)

/* Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

type counter struct {
	value int
	mutex sync.Mutex //Для первого способа
	ch    chan int   //Для второго способа
}

func main() {
	c := &counter{value: 0, mutex: sync.Mutex{}, ch: make(chan int, 1)}
	end := 50
	wg := sync.WaitGroup{}

	for i := 0; i < end; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment1(c)
		}()
	}
	wg.Wait()
	fmt.Println(c.value)

	c.value = 0

	for i := 0; i < end; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment2(c)
		}()
	}
	wg.Wait()
	close(c.ch)
	fmt.Println(c.value)
}

func increment1(c *counter) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func increment2(c *counter) {
	c.ch <- 1
	defer func() {
		<-c.ch
	}()
	c.value++
}
