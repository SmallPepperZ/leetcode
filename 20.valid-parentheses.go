/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main

// @lc code=start
func indexOf(s []rune, str rune) int {
	for i, v := range s {
		if v == str {
			return i
		}
	}
	return -1
}

func isIn(s []rune, str rune) bool {
	index := indexOf(s, str)
	return index != -1
}

func isInStr(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func isValid(s string) bool {
	chars := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	openingChars := make([]rune, 0, 3)
	charsRev := map[rune]rune{}
	for opening, closing := range chars {
		openingChars = append(openingChars, opening)
		charsRev[closing] = opening
	}
	pairs := []string{}

	for _, char := range s {
		if isIn(openingChars, char) {
			pairs = append(pairs, string(char))
		} else {
			openingChar := charsRev[char]
			if isInStr(pairs, string(openingChar)) {
				added := false
				for i := len(pairs) - 1; i >= 0; i-- {
					if len(pairs[i]) == 1 {
						pairs[i] += string(char)
						added = true
						break
					}
				}
				if !added {
					return false
				}
			} else {
				return false
			}
		}
	}
	for _, pair := range pairs {
		if pair != "()" && pair != "[]" && pair != "{}" {
			return false
		}
	}
	return true
}

// @lc code=end
