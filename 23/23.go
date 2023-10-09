package main

import "fmt"

//Удалить i-ый элемент из слайса

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	nums2 := []int{-41, -45, 11, 23, 35, 64, 34}
	i := 5
	fmt.Println(deleteNum1(nums1, i))
	fmt.Println(deleteNum2(nums2, i))
}

func deleteNum1(n []int, i int) []int {
	n = append(n[:i], n[i+1:]...)
	return n
}

func deleteNum2(n []int, i int) []int {
	copy(n[i:], n[i+1:])
	n[len(n)-1] = 0
	n = n[:len(n)-1]
	return n
}
