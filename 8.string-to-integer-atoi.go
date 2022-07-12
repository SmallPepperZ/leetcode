/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("20000000000000000000"))
}

// @lc code=start
var ints = map[rune]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'0': 0,
}

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	positive := true
	if s[0] == '-' {
		positive = false
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	outDigits := []int{}
	for _, char := range s {
		if integer, ok := ints[char]; ok {
			outDigits = append(outDigits, integer)
		} else {
			break
		}
	}
	out := float64(0)
	for i, digit := range outDigits {
		power := math.Pow10(len(outDigits) - i - 1)
		out += float64(digit) * power
	}

	if !positive {
		out = -out
	}
	if out > math.MaxInt32-1 {
		out = math.MaxInt32
	} else if out < math.MinInt32 {
		out = math.MinInt32
	}
	return int(out)
}

func checkLimit(num int64) (int64, bool) {

	if num > math.MaxInt32-1 {
		return math.MaxInt32, true
	}
	return num, false
}

// @lc code=end
