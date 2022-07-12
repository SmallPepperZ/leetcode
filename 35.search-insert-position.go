/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package main

import (
	"fmt"
)

// @lc code=start
func searchInsert(nums []int, target int) int {
	fmt.Println(nums)
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	center := len(nums) / 2
	if nums[center] == target {
		return center
	}
	if center != len(nums)-1 && nums[center+1] == target {
		return center + 1
	}
	if center != 0 && nums[center-1] == target {
		return center - 1
	}
	if nums[center] > target {
		if nums[center-1] < target {
			return center
		}
		return searchInsert(nums[:center], target)
	}
	if nums[center] < target {
		if nums[center+1] > target {
			return center + 1
		}
		newNums := nums[center:]
		return searchInsert(newNums, target) + len(nums) - len(newNums)
	}
	return -1
}

// @lc code=end
