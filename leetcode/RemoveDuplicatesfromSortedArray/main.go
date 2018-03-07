// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package main

import (
	"fmt"
)

// func removeDuplicates(nums []int) int {
// 	n := 0
// 	for i := 0; i < len(nums); i++ {
// 		if n == 0 || nums[i] != nums[n-1] {
// 			nums[n] = nums[i]
// 			n++
// 		}
// 	}
// 	return n
// }

// RemoveDuplicates just works as expected
func RemoveDuplicates(nums []int) int {
	b := nums[:0]
	for i, x := range nums {
		if i == 0 || x > nums[i-1] {
			b = append(b, x)
		}
	}
	return len(b)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Printf("%v\n", RemoveDuplicates(nums))
}
