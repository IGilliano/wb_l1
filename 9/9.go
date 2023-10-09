package main

import (
	"fmt"
)

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/

func main() {
	n := []int{2, 4, 6, 8, 10}
	nums := make(chan int)
	results := make(chan int)

	go func() {
		for i := range n {
			j := i
			nums <- n[j]
		}
		close(nums)
	}()

	go func() {
		for num := range nums {
			results <- num * 2
		}
		close(results)
	}()

	for num := range results {
		fmt.Println(num)
	}
}
