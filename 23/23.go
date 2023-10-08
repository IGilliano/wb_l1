package main

import "fmt"

//Удалить i-ый элемент из слайса

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i := 5
	copy(nums[i:], nums[i+1:])
	nums[len(nums)-1] = 0
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
}
