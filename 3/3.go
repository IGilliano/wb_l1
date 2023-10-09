package main

import (
	"fmt"
	"sync"
)

/* Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений. */

func main() {
	nums := []int{2, 4, 6, 8, 10}
	msg := make(chan int, len(nums))
	var sum int
	wg := sync.WaitGroup{}

	for i := range nums {
		wg.Add(1)
		num := nums[i]
		go func() {
			defer wg.Done()
			sum += num * num
		}()
	}
	wg.Wait()
	close(msg)

	for i := range msg {
		sum += i
	}
	fmt.Println(sum)
}
