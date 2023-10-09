package main

import (
	"fmt"
	"log"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	num := 123456
	fmt.Printf("Before: %b\n", num)
	var bit, set int
	fmt.Printf("What bit you want to format?\n")
	_, err := fmt.Scan(&bit)
	if err != nil {
		log.Fatal(err.Error())
	}

	if bit <= 0 && bit > 64 {
		log.Fatal("Incorrect")
	}

	fmt.Printf("1 or 0?\n")
	_, err = fmt.Scan(&set)
	if err != nil {
		log.Fatal(err.Error())
	}

	if set != 0 && set != 1 {
		log.Fatal("Incorrect")
	}

	setBit(num, bit, set)
}

func setBit(num int, bit int, set int) {
	if set == 0 {
		num = num &^ (1 << (bit - 1))
	} else {
		num = num | (1 << (bit - 1))
	}
	fmt.Printf("After:  %b\n", num)
}
