package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
)

/**
https://neetcode.io/problems/top-k-elements-in-list
Top K Frequent Elements
Given an integer array nums and an integer k, return the k most frequent elements within the array.
**/

func topKFrequent(nums []int, k int) []int {
	var kMostFrequent []int
	// map of numbers and their frequencies
	var kMap = make(map[int]int)

	for _, num := range nums {
		_, exists := kMap[num]
		if exists {
			kMap[num] += 1
		} else {
			kMap[num] = 1
		}
	}

	var allFrequencies []int
	for _, value := range kMap {
		allFrequencies = append(allFrequencies, value)
	}

	slices.SortFunc(allFrequencies, func(a, b int) int {
		return b - a
	})

	// get top K frequency number
	var topKFrequencies = allFrequencies[:k]

	// get the numbers with frequencies from topKFrequencies
	for key, num := range kMap {
		if slices.Contains(topKFrequencies, num) {
			kMostFrequent = append(kMostFrequent, key)
		}
	}

	return kMostFrequent
}

// bucket sort
func topKFrequent2(nums []int, k int) []int {
	var numsLen = len(nums)
	var kMostFrequent []int
	var bucket = make([][]int, numsLen, numsLen)

	var min = slices.Min(nums)
	var max = slices.Max(nums)
	var numRange = max - min
	var bucketCount = numsLen
	var bucketSize = int(math.Ceil(float64((numRange + 1)) / float64(bucketCount)))

	for _, num := range nums {
		var bucketIndex = (num - min) / bucketSize
		bucket[bucketIndex] = append(bucket[bucketIndex], num)
	}

	var bucketLengthsMap = make(map[int]int)  // frequency as key, number as value
	var bucketLengthsMap2 = make(map[int]int) // frequency as key, number as value
	maps.Copy(bucketLengthsMap2, bucketLengthsMap)
	var bucketLengths []int // contains the lengths of each buckets
	for _, b := range bucket {
		if len(b) > 0 {
			bucketLengthsMap[len(b)] = b[0]
		}
		bucketLengths = append(bucketLengths, len(b))
	}

	// fmt.Println("bucketlength2", bucketLengthsMap, bucketLengths)
	for range k {
		var m = slices.Max(bucketLengths)
		// fmt.Println("ranging k", bucketLengthsMap, m)
		kMostFrequent = append(kMostFrequent, bucketLengthsMap[m])

		var tempCont []int
		for _, num := range bucketLengths {
			if num != m {
				tempCont = append(tempCont, num)
			}
		}
		bucketLengths = tempCont
	}

	// fmt.Println("buckets", bucket, bucketCount, bucketSize, bucket, bucketLengthsMap, kMostFrequent)
	return kMostFrequent
}

// bucket sort with hint
func topKFrequent3(nums []int, k int) []int {
	var topKFrequent []int
	var frequencies = make(map[int]int)

	for _, num := range nums {
		var _, exists = frequencies[num]
		if exists {
			frequencies[num] += 1
		} else {
			frequencies[num] = 1
		}
	}

	// an array whose index represents the frequency
	// the +1 is for the edge case where nums has same number.
	var buckets = make([][]int, len(nums)+1)
	// var buckets [][]int

	for num, frequency := range frequencies {
		// fmt.Println("buckets", buckets)
		// fmt.Println("freq", frequency, buckets[frequency])
		buckets[frequency] = append(buckets[frequency], num)
	}

	for i := range buckets {
		topKFrequent = append(topKFrequent, buckets[(len(buckets))-i-1]...)
		fmt.Println("topk", buckets[(len(buckets))-i-1], topKFrequent)

		if len(topKFrequent) >= k {
			break
		}
	}

	return topKFrequent[:k]
}

func main() {
	// var kMostFrequent = topKFrequent([]int{1, 2, 2, 3, 3, 4, 4, 4, 4}, 2)
	// var kMostFrequent = topKFrequent([]int{7, 7}, 1)
	// var kMostFrequent = topKFrequent2([]int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 5, 5, 5, 5, 5, 5}, 1)
	var kMostFrequent = topKFrequent3([]int{1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4}, 2)
	fmt.Println("kMostFrequent", kMostFrequent)
}
