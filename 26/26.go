package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false

*/

func main() {
	strings := []string{"abcd", "abCdefAaf", "aabcd"}

	for i := range strings {
		fmt.Println(unicCheck(strings[i]))
	}
}

func unicCheck(s string) bool {
	symbols := make(map[rune]int)
	s = strings.ToLower(s)
	for _, v := range s {
		symbols[v] += 1
	}

	for k := range symbols {
		if symbols[k] > 1 {
			return false
		}
	}
	return true
}
