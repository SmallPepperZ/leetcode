/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

// @lc code=start


func romanToInt(s string) int {
	lookup := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	values := []int{}
	sum := 0
    for _, char := range s {
		value := lookup[char]
		values = append(values, value)
	}
	for i, value := range values {
		if i == len(values)-1 || value >= values[i+1] {
			sum += value
		} else {
			sum -= value
		}
	}
	return sum
}
// @lc code=end

