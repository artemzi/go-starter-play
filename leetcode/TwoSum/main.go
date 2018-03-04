// Two Sum code problem
// https://leetcode.com/problems/two-sum/description/
package main

import "fmt"

var (
	nums   = []int{2, 5, 5, 11}
	target = 10
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum(nums, target))
}
