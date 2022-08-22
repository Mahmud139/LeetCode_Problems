package main

import "fmt"

func main() {
	test := []int{3,2,4,5,6}
	target := 9
    fmt.Println(twoSum(test, target))
}

func twoSum(nums []int, target int) []int {
    result := map[int]int{}
    for idex, value := range nums {
        remaining := target - value
        if _, ok := result[remaining]; ok {
            return []int{result[remaining], idex}
        } else {
            result[value] = idex
        }
    }
    return []int{}
}