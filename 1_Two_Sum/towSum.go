package main

import "fmt"

func main() {
	test := []int{3,2,4,5,6}
	target := 9
	fmt.Println(twoSum(test, target))
}

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
    }
    return nil
}
