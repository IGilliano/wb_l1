package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type Map struct {
	mutex sync.Mutex
	m     map[int]string
}

func main() {
	test := Map{mutex: sync.Mutex{}, m: make(map[int]string)}
	values := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven"}
	wg := sync.WaitGroup{}

	for i := range values {
		j := i
		wg.Add(1)
		go func() {
			test.mutex.Lock()
			defer test.mutex.Unlock()
			test.m[j] = values[j]
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(test.m)
}
