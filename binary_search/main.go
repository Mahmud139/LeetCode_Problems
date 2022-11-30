//binary search only working on sorted array/slice.
package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(binarySearch(a, 6))
}

func binarySearch(a []int, x int) int {
	left := 0
	right := len(a) - 1
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == x {
			return mid
		} else if a[mid] < x {
			left = mid + 1
		} else {
			right = mid -1
		}
	}
	return -1
}

// time complexity of this algorithm : O(log n) or O(log2 n)
// space complexity of this algorithm : O(1)