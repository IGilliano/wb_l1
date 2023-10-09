package main

import (
	"fmt"
	"log"
)

/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/

func main() {
	var s string
	fmt.Println("Input string\n")
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}
