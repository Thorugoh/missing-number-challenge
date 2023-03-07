package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result := findMissingNumber(arr)

	fmt.Println(strconv.Itoa(result))
}

func findMissingNumber(arr []int) int {
	arrLen := len(arr)
	expectedSum := arrLen * (arrLen + 1) / 2
	arrSum := 0

	for _, num := range arr {
		arrSum += num
	}

	return expectedSum - arrSum
}
