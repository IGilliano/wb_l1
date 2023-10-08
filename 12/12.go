package main

import "fmt"

/*Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.*/

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)
	results := make([]string, 0)

	for _, v := range input {
		m[v]++
	}

	for k := range m {
		results = append(results, k)
	}
	fmt.Println(results)
}
