package main

import "fmt"

/*
https://neetcode.io/problems/missing-number
*/

func findMissingNumber(nums []int) int {
	numMap := map[int]int{}

	// range of elements is [0 , len(nums)]
	// lay out all expected numbers in a map
	for i := range len(nums) {
		numMap[i] = 0
	}

	for _, n := range nums {
		_, exists := numMap[n]
		if exists {
			numMap[n] = 1
		}
	}

	for i := range numMap {
		if numMap[i] == 0 {
			return i
		}
	}

	return 0
}

func findMissingNumberOptimized(nums []int) int {
	nl := len(nums)
	expectedSum := (nl * (nl + 1)) / 2
	actualSum := 0

	for _, n := range nums {
		actualSum += n
	}

	return expectedSum - actualSum
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3}, want: 0},
		{nums: []int{0, 2}, want: 1},
		{nums: []int{5, 2, 6, 0, 1, 3}, want: 4},
	}

	for i, t := range testCases {
		// got := findMissingNumber(t.nums)
		got := findMissingNumberOptimized(t.nums)
		if t.want == got {
			fmt.Printf("Test %d: got %d, want %d passed\n", i, got, t.want)
		} else {
			fmt.Printf("Test %d: got %d, want %d failed\n", i, got, t.want)
		}
		fmt.Println("------------------------------------")
	}
}
