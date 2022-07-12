/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package main

// @lc code=start
var mapping map[rune][]string = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	output := make([]string, 1)
	for _, digit := range digits {
		letters := mapping[digit]
		output = getCombinations(output, letters)
	}
	if len(output) == 1 {
		return []string{}
	}
	return output
}

func getCombinations(a []string, b []string) []string {
	output := make([]string, 0, len(a)*len(b))
	for _, valA := range a {
		for _, valB := range b {
			output = append(output, valA+valB)
		}
	}
	return output
}

// @lc code=end
