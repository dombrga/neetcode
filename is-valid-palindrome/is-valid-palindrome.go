/*
https://neetcode.io/problems/is-palindrome
Given a string s, return true if it is a palindrome, otherwise return false.

A palindrome is a string that reads the same forward and backward.
It is also case-insensitive and ignores all non-alphanumeric characters.
*/
package main

import (
	"fmt"
	"strings"
)

type testCase struct {
	s            string
	isPalindrome bool
}

// first solution
func isValidPalindrome(s string) bool {
	var fs = strings.ToLower(strip(s))
	var sLen = len(fs)
	// integer division, ie 5/2 = 2
	var sLenHalf = sLen / 2

	for i := 0; i < sLenHalf; i++ {
		if !(fs[i] == fs[sLen-1-i]) {
			return false
		}
	}

	return true
}

// two pointer algo
func isValidPalindrome2(s string) bool {
	var left = 0
	var right = len(s) - 1

	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		// for {
		// 	if !isAlphaNumeric(s[left]) {
		// 		left++
		// 	}
		// 	if !isAlphaNumeric(s[right]) {
		// 		right--
		// 	}

		// 	if isAlphaNumeric(s[left]) && isAlphaNumeric(s[right]) {
		// 		break
		// 	}
		// }

		// fmt.Printf("letter: %c %c %d %d\n", s[left], s[right], left, right)
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left++
		right--
	}
	return true
}

// based on https://stackoverflow.com/questions/54461423/efficient-way-to-remove-all-non-alphanumeric-characters-from-large-text
// removes non-alphanumeric characters and space in s
func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			// b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func isAlphaNumeric(s byte) bool {
	return ('a' <= s && s <= 'z') ||
		('A' <= s && s <= 'Z') ||
		('0' <= s && s <= '9')
}

func main() {
	var testCases = []testCase{
		{"racecar", true},
		{"Was it a car ?or a cat I saw?", true},
		{"tab a cat", false},
		{"tab zxcsdawe?a cat", false},
	}

	for _, c := range testCases {
		// var isValidPalindrome = isValidPalindrome(c.s)
		var isValidPalindrome = isValidPalindrome2(c.s)
		// fmt.Printf("\"%s\" is valid palindrome: %t, %t\n", c.s, isValidPalindrome, isValidPalindrome == c.isPalindrome)
		if isValidPalindrome == c.isPalindrome {
			if isValidPalindrome {
				fmt.Printf("Test passed: \"%s\" is a palindrome\n", c.s)
			} else {
				fmt.Printf("Test passed: \"%s\" is not a palindrome\n", c.s)
			}
		} else {
			fmt.Printf("Test failed: \"%s\"\n", c.s)
		}
	}
}
