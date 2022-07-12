/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package main
// @lc code=start
func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i, char := range strs[0] {
		if isCommon(strs, char, i) {
			prefix += string(char)
		} else {
			break
		}
	}
	return prefix
}

func isCommon(strs []string, char rune, i int) bool {
	for _, str := range strs {
		if i >= len(str) {
			return false
		}
		if rune(str[i]) != char {
			return false
		}
	}
	return true
}
// @lc code=end

