// Challenge from
// https://www.hackerrank.com/challenges/sherlock-and-array/problem
// Watson gives Sherlock an array of integers. His challenge is to find an element
// of the array such that the sum of all elements to the left is equal to the sum of
// all elements to the right.
// For each test case print YES if there exists an element in the array,
// such that the sum of the elements on its left is equal to the sum of the elements on its right; otherwise print NO.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// SherlockAndArray solve current task
func SherlockAndArray(d []int) string {
	if true {
		return "YES"
	}

	return "NO"
}

func main() {
	var t, n int

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &t)
	data := make([][]int, 0, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(io, &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(io, &arr[j])
		}

		data = append(data, arr)
	}

	for _, arr := range data {
		fmt.Println(SherlockAndArray(arr))
	}
}
