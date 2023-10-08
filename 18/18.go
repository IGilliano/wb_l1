package main

import (
	"fmt"
	"os"
	"sync"
)

/* Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика. */

type counter struct {
	value int
	mutex sync.Mutex //Для первого способа
	ch    chan int   //Для второго способа
}

func main() {
	c := counter{value: 0, mutex: sync.Mutex{}, ch: make(chan int, 1)}
	end := 50
	wg := sync.WaitGroup{}

	for i := 0; i < end; i++ {
		wg.Add(1)
		go increment(&c, &wg)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, c.value)

	c.value = 0

	for i := 0; i < end; i++ {
		wg.Add(1)
		go increment2(&c, &wg)
	}
	wg.Wait()
	close(c.ch)
	fmt.Fprintln(os.Stdout, c.value)
}

func increment(c *counter, wg *sync.WaitGroup) {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
	wg.Done()
}

func increment2(c *counter, wg *sync.WaitGroup) {
	c.ch <- 1
	c.value++
	<-c.ch
	wg.Done()
}
