// Package main solve:
// https://leetcode.com/problems/search-insert-position/description/
// Given a sorted array and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

// You may assume no duplicates in the array.
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, digit := range nums {
		if digit < target {
			continue
		}
		return i
	}
	return len(nums)
}

func main() {
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 7)) // 4
	fmt.Printf("%v\n", searchInsert([]int{1, 3, 5, 6}, 0)) // 0
}
