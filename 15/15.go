package main

import (
	"fmt"
	"strings"
)

/*

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func createHugeString(num int) string {
	str := strings.Repeat("界", num)
	return str
}

/*Первая и главная проблема: создавая срез изначальной строки мы сохраняем указатель на нее,
из-за чего в памяти хранится и срез,и огромная строка.
Чтобы избежать этого и позволить сборщику мусора убрать ненужную строку, нужно использовать функцию Copy */

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
	/*Строка в Go представлена как слайс байтов, поэтому если она содержит символы unicode,
	то при подобной индексации часть строки потеряется. Поэтому далее я преобразую строку в слайс рун.*/
	hugeRuneString := []rune(v)
	hugeRuneStringCopy := make([]rune, 100)
	copy(hugeRuneStringCopy, hugeRuneString)
	justString = string(hugeRuneStringCopy)
	fmt.Println(justString)
}
func main() {
	someFunc()
}
