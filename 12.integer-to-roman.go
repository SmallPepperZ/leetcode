/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(40))
	fmt.Println(intToRoman(90))
	fmt.Println(intToRoman(400))
	fmt.Println(intToRoman(900))
	fmt.Println(intToRoman(941))
}

// @lc code=start
var numeralMap = map[int]rune{
	1000: 'M',
	500:  'D',
	100:  'C',
	50:   'L',
	10:   'X',
	5:    'V',
	1:    'I',
}

func intToRoman(num int) string {

	keys := make([]int, 0, len(numeralMap))
	for key := range numeralMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	output := []rune{}
	for i, key := range keys {
		if key <= num {
			count := num / key
			if count == 4 {
				if len(output) != 0 && output[len(output)-1] == numeralMap[keys[i-1]] {
					output = append(output[:len(output)-1], numeralMap[key], numeralMap[keys[i-2]])
				} else {
					output = append(output, numeralMap[key], numeralMap[keys[i-1]])
				}

			} else {
				for i := 0; i < count; i++ {
					output = append(output, numeralMap[key])
				}

			}
			num -= key * count
		}
	}
	return string(output)
}

// @lc code=end
