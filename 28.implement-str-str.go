/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
package main

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i, char := range haystack {
		match := false
		if char == rune(needle[0]) {
			match = true
			for i1, char1 := range needle {
				if i+i1 >= len(haystack) {
					return -1
				}
				if rune(haystack[i+i1]) != char1 {
					match = false
				}
			}

		}
		if match {
			return i
		}
	}
	return -1
}

// @lc code=end
