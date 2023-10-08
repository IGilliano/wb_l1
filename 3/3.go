package main

import (
	"fmt"
	"sync"
)

/* Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений. */

func main() {
	nums := []int{2, 4, 6, 8, 10}
	msg := make(chan int)
	var sum int
	wg := sync.WaitGroup{}

	for i := range nums {
		wg.Add(1)
		num := nums[i]
		go func() {
			msg <- num * num
		}()
		wg.Done()
		sum += <-msg
	}
	wg.Wait()
	close(msg)
	fmt.Println(sum)
}
