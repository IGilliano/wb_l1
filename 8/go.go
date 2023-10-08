package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	num := 123456
	fmt.Printf("Before: %b\n", num)
	bits := []int{2, 11, 23}
	set := []int{1, 0, 1}

	for i := 0; i < 3; i++ {
		setBit(num, bits[i], set[i])
	}
}

func setBit(num int, bit int, set int) {
	if set == 0 {
		num = num &^ (1 << (bit - 1))
	} else {
		num = num | (1 << (bit - 1))
	}
	fmt.Printf("After:  %b\n", num)
}
