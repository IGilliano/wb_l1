package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/* Реализовать быструю сортировку массива (quicksort) встроенными методами языка. */

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := rand.Int() % len(nums)
	left, right := 0, len(nums)-1
	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	quicksort(nums[:left])
	quicksort(nums[left+1:])

	return nums
}

func main() {
	nums := []int{15, 25, -55, 66, 865, -2}
	fmt.Println(quicksort(nums))
	//На случай, если вопрос с подвохом:
	sort.Ints(nums)
	fmt.Println(nums)
}
