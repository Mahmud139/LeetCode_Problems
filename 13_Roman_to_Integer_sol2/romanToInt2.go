package main

import "fmt"

func main() {
	romanChar := "MCMXCIV"
	fmt.Println(romanToInt(romanChar))
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	num, lastInt, sum := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		num = roman[string(char)]
		if num >= lastInt {
			lastInt = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}