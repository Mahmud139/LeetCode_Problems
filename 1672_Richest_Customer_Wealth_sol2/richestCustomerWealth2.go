package main

import (
	"fmt"
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
	wealthCusMoney, total := 0, 0
	for i := 0; i < len(accounts); i++ {
		total = 0
		for j := 0; j < len(accounts[i]); j++ {
			total = total + accounts[i][j]
		}
		// for _, value := range accounts[i] {
		// 	total = total + value
		// }
		if wealthCusMoney < total {
			wealthCusMoney = total
		}
	}
	return wealthCusMoney
}