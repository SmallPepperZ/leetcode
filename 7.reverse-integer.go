/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-2143847412))
	fmt.Println(reverse(1563847412))
	fmt.Println(math.MaxInt32)
}

// @lc code=start
func reverse(x int) int {
	inputString := strconv.Itoa(x)
	multiplier := 1
	if inputString[0] == '-' {
		multiplier = -1
		inputString = inputString[1:]
	}
	output := 0
	for i, num := range inputString {
		numScaled, _ := strconv.Atoi(string(num))
		for i1 := 0; i1 < i; i1++ {
			if numScaled == 0 {
				break
			} else if (math.MaxInt32-output)/numScaled >= 10 {
				numScaled *= 10
			} else {
				return 0
			}
		}
		output += numScaled
	}
	return output * multiplier
}

// @lc code=end
