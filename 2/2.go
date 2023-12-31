package main

import (
	"fmt"
	"sync"
)

/* Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout. */

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for i := range nums {
		wg.Add(1)
		num := nums[i]
		go func() {
			defer wg.Done()
			sq := num * num
			fmt.Println(sq)
		}()
	}
	wg.Wait()
}
