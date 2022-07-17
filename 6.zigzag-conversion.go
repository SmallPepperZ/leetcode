/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */
package main

import "fmt"

func main() {
	fmt.Println(convert("AB", 1))
	fmt.Println(convert("PAYPALISHIRING", 3))
}

// @lc code=start
func convert(s string, numRows int) string {
	rows := make([][]rune, numRows)
	rIndex := -1
	increment := 1
	if numRows == 1 {
		return s
	}
	for _, char := range s {
		rIndex += increment

		if rIndex == numRows {
			rIndex = numRows - 2
			increment = -1
		}
		if rIndex == -1 {
			rIndex = 1
			increment = 1
		}
		rows[rIndex] = append(rows[rIndex], char)

	}
	output := []rune{}
	for _, row := range rows {
		output = append(output, row...)
	}
	return string(output)
}

// @lc code=end
