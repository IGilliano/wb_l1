package main

import (
	"fmt"
	"strings"
)

/*

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

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
	str := strings.Repeat("o", num)
	str = str + "a"
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v)
	justString = v[:100]
	fmt.Println(justString)

}

func main() {
	someFunc()
}
