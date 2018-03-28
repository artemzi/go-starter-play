package main

import (
	"fmt"
)

func insertionsort2(data []int64) {
	l := len(data)

	for i := 1; i < l; i++ {
		x := data[i]
		j := i

		for j > 0 && data[j-1] > x {
			data[j] = data[j-1]
			j--
		}
		data[j] = x

		for i := 0; i < l; i++ {
			fmt.Printf("%v ", data[i])
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	insertionsort2(data)
}
