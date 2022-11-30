package main


func firstMissingPositive(nums []int) int {
	// Linear time, and constant extra space.
	// Unrealistic solution, because it loses all the values initially <= 0:
	// the negative numbers already in the array can't be recovered without extra space for remembering them.
	// It's fine if we don't care about the values, or all the values are positive.

	limit := len(nums)
	// For correctness, ensure no value is negative,
	// making all of them between 1 and limit + 1, both included.
	for i, num := range nums {
		if num <= 0 || num > limit {
			nums[i] = limit + 1
		}
	}

	// Mark the appearance of positive number x (1, 2, ..., limit)
	// by negating the value at position x-1 (0, 1, ..., limit-1)
	for _, num := range nums {
		absNum := num
		if absNum < 0 {
			absNum = -absNum
		}
		if absNum > limit {
			continue
		}
		appearanceMarker := nums[absNum-1]
		if appearanceMarker > 0 {
			nums[absNum-1] = -appearanceMarker
		}
	}

	// Look for the first unmarked position, corresponding to number position+1
	position := 0
	for position < limit {
		if nums[position] > 0 {
			break
		}
		position++
	}
	return position + 1
}