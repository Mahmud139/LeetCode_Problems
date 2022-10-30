package main

import (
	"fmt"
)

func main() {
	path := "UDDDUDUU"

	fmt.Println(countingValleys(path))
}

func countingValleys(path string) int32 {
	lavel, valley := 0, 0

	for i := 0; i < len(path); i++ {
		if path[i] == 'U' {
			lavel++
		} else if path[i] == 'D' {
			if lavel == 0 {
				valley++
			}
			lavel--
		}
	}

	return int32(valley)
}