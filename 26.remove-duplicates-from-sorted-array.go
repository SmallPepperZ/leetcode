/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	uniqueCount := 0
	for i, num := range nums {
		isDuplicate := false
		for _, num1 := range nums[:i] {
			if num == num1 {
				isDuplicate = true
			}
		}
		if !isDuplicate {
			nums[uniqueCount] = num
			uniqueCount += 1
		}
	}
	return uniqueCount
}

// @lc code=end
