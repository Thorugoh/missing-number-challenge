package main

import "testing"

func TestMissingNumber(t *testing.T) {

	result := findMissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})

	if result != 8 {
		t.Errorf("FAILED. Expected %d, got %d \n", 8, result)
	}
}

func TestMissingNumber2(t *testing.T) {

	result := findMissingNumber([]int{0, 1, 2, 3, 5})

	if result != 4 {
		t.Errorf("FAILED. Expected %d, got %d \n", 4, result)
	}
}
