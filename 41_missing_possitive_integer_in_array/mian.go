package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,20}
	c := solution(a)
	fmt.Println("result: ", c)
}

func solution(a []int) int {
	ex := make([]int, 0)
	
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			continue
		} else {
			ex = append(ex, a[i])
		}
	}
	if len(ex) == 0 {
		return 1
	}

	check := make(map[int]int)
	res := make([]int, 0)
	
	for _, val := range ex {
		check[val] = 1
	}

	for letter := range check {
		res = append(res, letter)
	}
	
	sort.Ints(res)
	fmt.Println(res)
	for i := 1; i <= len(res); i++ {
		if i == res[i-1] {
			if i == len(res) {
				return i+1
			}
			continue
		} else {
			return i
		}
	}
	return 0
}