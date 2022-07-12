/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

package main

import (
	"strconv"
)

// @lc code=start

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []rune
	for {
		s1 = append(s1, rune(strconv.Itoa(l1.Val)[0]))
		if l1.Next == nil {
			break
		}
		l1 = l1.Next

	}
	for {
		s2 = append(s2, rune(strconv.Itoa(l2.Val)[0]))
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}
	sum := add(s1, s2)
	// n1, _ := strconv.ParseFloat(s1, 64)
	// n2, _ := strconv.ParseFloat(s2, 64)

	// sum := strconv.FormatFloat(n1+n2, 'f', 0, 64)
	// fmt.Println(s1, s2, n1, n2, sum)
	outputHead := &ListNode{}
	output := outputHead
	for _, val := range sum {
		valInt, _ := strconv.Atoi(string(val))
		output.Next = &ListNode{valInt, nil}
		output = output.Next
	}
	return outputHead.Next
}

func reverse(s []rune) []rune {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}
func add(s1, s2 []rune) []rune {
	out := []rune{}
	carry := false
	n1 := s1
	n2 := s2
	if len(s1) < len(s2) {
		n1 = s2
		n2 = s1
	}
	for i, r1 := range n1 {
		v1, _ := strconv.Atoi(string(r1))
		v2 := 0
		if i < len(n2) {
			v2, _ = strconv.Atoi(string(n2[i]))
		}
		sum := v1 + v2
		if carry {
			sum += 1
			carry = false
		}

		if sum > 9 {
			carry = true
			sum -= 10
		}

		out = append([]rune{rune(strconv.Itoa(sum)[0])}, out...)
	}
	if carry {
		out = append([]rune{'1'}, out...)
	}
	return reverse(out)
}

// @lc code=end

// 942
// 9465
// 1884

// 10407
