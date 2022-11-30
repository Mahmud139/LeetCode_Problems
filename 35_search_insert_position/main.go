package main

import "fmt"

func main() {
	n := []int{1,3,5,6}
	fmt.Println(searchInsert(n, 7))
}

func searchInsert(nums []int, target int) int {
    left := 0
	right := len(nums)-1
	for left <= right {
		mid := (left + right) /2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid -1
		}
	}
	return left
}