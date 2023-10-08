package main

import "fmt"

/*Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.*/

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	results := make(map[int][]float32)

	for _, v := range temp {
		key := int(v) / 10 * 10
		results[key] = append(results[key], v)
	}

	fmt.Println(results)
}
