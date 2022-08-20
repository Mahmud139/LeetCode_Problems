package main

import (
	"fmt"
	"sort"
)

func main() {
	test := [][]int{
		{1, 2, 3},
		{3, 1, 1},
		{10, 1, 1},
	}
	fmt.Println(maximumWealth(test))
}

func maximumWealth(accounts [][]int) int {
	var wealthCus []int
	for i := 0; i < len(accounts); i++ {
		var sum int = 0
		for _, value := range accounts[i] {
			sum = sum + value
		}
		wealthCus = append(wealthCus, sum)
	}
	sort.Ints(wealthCus)
	return wealthCus[len(wealthCus) - 1]
}