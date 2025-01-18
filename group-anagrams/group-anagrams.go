package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
https://neetcode.io/problems/anagram-groups
Group Anagrams
Given an array of strings strs, group all anagrams together into sublists.
You may return the output in any order.

An anagram is a string that contains the exact same characters as another string,
but the order of the characters can be different.
**/

func groupAnagrams(strs []string) [][]string {
	var groupedAnagrams [][]string
	var anagramMap = make(map[string][]string)

	for _, word := range strs {
		var sorted = sortString(word)
		var _, exists = anagramMap[sorted]
		if !exists {
			anagramMap[sorted] = []string{word}
		} else {
			anagramMap[sorted] = append(anagramMap[sorted], word)
		}
	}

	for _, value := range anagramMap {
		groupedAnagrams = append(groupedAnagrams, value)
	}

	return groupedAnagrams
}

func groupAnagrams2(strs []string) [][]string {
	letterMap := make(map[[26]int][]string)
	fmt.Println("letterMap", letterMap)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		letterMap[count] = append(letterMap[count], s)
		fmt.Println("count", count, letterMap[count])
	}

	var result [][]string
	for _, group := range letterMap {
		result = append(result, group)
	}
	return result
}

func sortString(w string) string {
	var s = strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	// var groupedAnagrams = groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})
	// fmt.Println("group anagrams", groupedAnagrams, 'c'-'a', 'c', 'a')
	var groupedAnagrams2 = groupAnagrams2([]string{"act", "pots", "tops", "cat", "stop", "hat"})
	fmt.Println("group anagrams", groupedAnagrams2, 'c'-'a')
}
