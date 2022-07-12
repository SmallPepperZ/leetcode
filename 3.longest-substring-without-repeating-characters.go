/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package main



// @lc code=start
func lengthOfLongestSubstring(s string) int {
	substrings := []string{}
	for i, char := range s {
		substring := []rune{char}
		// if i == len(s)-1 {
		// 	substrings = append(substrings, string(char))
		// }
		for _, char2 := range s[i+1:] {
			if !isElementExist(substring, char2) {
				substring = append(substring, char2)
			} else {
				break
			}
		}
		substrings = append(substrings, string(substring))
	}
	max := 0
	for _, substring := range substrings {
		if len(substring) > max {
			max = len(substring)
		}
	}
	return max
}

func isElementExist(s []rune, str rune) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// @lc code=end
