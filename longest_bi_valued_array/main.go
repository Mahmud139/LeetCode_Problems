package main

import "fmt"

import "sort"

func solution(A []int) int {
	if len(A) == 2 {
		return 2
	}

	// sum := 0
	// sort.Ints(A)
	check := make(map[int]int)
	res := make([]int, 0)

	for _, val := range A {
		check[val] += 1
	}
	fmt.Println(check)

	for _, v := range check {
		res = append(res, v)
	}

	sort.Ints(res)
	fmt.Println(res)
	// for i, _ := range check {
	// 	if check[i] <= 1 {
	// 		continue
	// 	}
	// 	sum += check[i]
	// }

	// if sum == 2 {
	// 	return sum+1
	// }
	return res[len(res)-1]+res[len(res)-2]
}

func main() {
	A := []int{0,2,2,4,3,4,3,2,3,12} //0,2,2,2,3,3,4,4,12
	
	fmt.Println(solution(A))
}
