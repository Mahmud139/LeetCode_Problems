package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	res := linearSearch(a, 5)
	fmt.Println(res)
}

func linearSearch(a []int, x int) int {
	for i := range a {
		if a[i] == x {
			return i
		}
	}
	return -1
}