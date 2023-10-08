package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 56, 25, 233, 33, 44, 16, -6, -40}
	sort.Ints(nums)
	fmt.Println(nums)
	x := 44 //Искомое число
	result := binarySearch(nums, x)
	if result == -1 {
		fmt.Printf("Num %d is not found", x)
	} else {
		fmt.Printf("Num %d has index %d", x, result)
	}
}

func binarySearch(nums []int, x int) int {
	result := -1 //Если число не будет найдено - функция вернет -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == x {
			result = mid //Corner case
			break
		} else if nums[mid] < x {
			left = mid + 1
		} else if nums[mid] > x {
			right = mid - 1
		}
	}
	return result
}
