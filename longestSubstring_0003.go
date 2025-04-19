package main

/*
Given a string s, find the length of the longest

without duplicate characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/
func lengthOfLongestSubstring(s string) int {
	start := 0
	max := 0
	m := make(map[rune]int)
	for i, char := range s {
		if lastIdx, found := m[char]; found && lastIdx >= start {
			start = lastIdx + 1
		}
		m[char] = i
		if max < i-start+1 {
			max = i - start + 1
		}
	}
	return max
}
