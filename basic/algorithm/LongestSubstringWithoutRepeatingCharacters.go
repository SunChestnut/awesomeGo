package main

import "fmt"

/**
 * [LeetCode 3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
 */

// 不能计算中文字符版
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 可以计算中文字符版
func longestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {

	fmt.Println(
		longestSubstring("abcabcbb"),
		longestSubstring("咿呀咿呀呦嘿"),
	)

}
