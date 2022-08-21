package main

import "fmt"

func main() {
	test := []int{3,2,3,4}
	target := 6
	fmt.Println(twoSum(test, target))
}

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        if nums[i] + nums[i+1] == target {
            return []int{i, i+1}
        }
    }
    return nil
}