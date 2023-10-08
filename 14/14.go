package main

import "fmt"

/* Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.*/

// 1 способ: использовать оператор switch:
func identifyType1(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("%d = int!\n", x)
	case string:
		fmt.Printf("%s = string!\n", x)
	case float64:
		fmt.Printf("%.2f = float!\n", x)
	case rune:
		fmt.Printf("%q = rune!\n", x)
	case bool:
		fmt.Printf("%t = bool!\n", x)
	}
}

// 2 способ: использовать Printf:
func identifyType2(x interface{}) {
	fmt.Printf("%v sure is %T!\n", x, x)
}

func main() {
	var input []interface{}
	input = append(input, 12)
	input = append(input, "Twelve")
	input = append(input, 12.2)
	input = append(input, '¿')
	input = append(input, true)

	for _, v := range input {
		identifyType1(v)
		identifyType2(v)
	}
}
